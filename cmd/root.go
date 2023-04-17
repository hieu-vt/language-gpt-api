package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/component/ginc"
	smdlw "github.com/viettranx/service-context/component/ginc/middleware"
	"github.com/viettranx/service-context/component/gormc"
	"github.com/viettranx/service-context/component/jwtc"
	"google.golang.org/grpc"
	"lang-gpt-api/common"
	"lang-gpt-api/composer"
	"lang-gpt-api/middleware"
	"lang-gpt-api/proto/pb"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("Language Microservices with GPT"),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompMySQL, "")),
		sctx.WithComponent(jwtc.NewJWT(common.KeyCompJWT)),
		sctx.WithComponent(NewConfig()),
	)
}

type RequestParams struct {
	Mess string `json:"mess"`
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		// Make some delay for DB ready (migration)
		// remove it if you already had your own DB
		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(common.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx))

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		router.GET("/gpt/:mess", func(c *gin.Context) {
			mess := c.Param("mess")

			data := DemoGptAPI(mess)
			c.JSON(http.StatusOK, gin.H{"data": data})
		})

		go StartGRPCServices(serviceCtx)

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

type RequestBody struct {
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
	MaxLength   int     `json:"max_length"`
}

type ResponseBody struct {
	GeneratedText string `json:"generated_text"`
}

func DemoGptAPI(mess string) *ResponseBody {
	apiKey := ""
	//apiUrl := "https://api-inference.huggingface.co/models/EleutherAI/gpt-neo-1.3B"

	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: mess,
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &ResponseBody{GeneratedText: resp.Choices[0].Message.Content}
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {

	userAPIService := composer.ComposeUserAPIService(serviceCtx)
	taskAPIService := composer.ComposeTaskAPIService(serviceCtx)
	authAPIService := composer.ComposeAuthAPIService(serviceCtx)

	requireAuthMdw := middleware.RequireAuth(composer.ComposeAuthRPCClient(serviceCtx))

	router.POST("/authenticate", authAPIService.LoginHdl())
	router.POST("/register", authAPIService.RegisterHdl())
	router.GET("/profile", requireAuthMdw, userAPIService.GetUserProfileHdl())

	tasks := router.Group("/tasks", requireAuthMdw)
	{
		tasks.GET("", taskAPIService.ListTaskHdl())
		tasks.POST("", taskAPIService.CreateTaskHdl())
		tasks.GET("/:task-id", taskAPIService.GetTaskHdl())
		tasks.PATCH("/:task-id", taskAPIService.UpdateTaskHdl())
		tasks.DELETE("/:task-id", taskAPIService.DeleteTaskHdl())
	}
}

func StartGRPCServices(serviceCtx sctx.ServiceContext) {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)
	logger := serviceCtx.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configComp.GetGRPCPort()))

	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("GRPC Server is listening on %d ...\n", configComp.GetGRPCPort())

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, composer.ComposeUserGRPCService(serviceCtx))
	pb.RegisterAuthServiceServer(s, composer.ComposeAuthGRPCService(serviceCtx))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
