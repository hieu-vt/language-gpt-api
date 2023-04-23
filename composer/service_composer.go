package composer

import (
	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"lang-gpt-api/common"
	"lang-gpt-api/proto/pb"
	authBusiness "lang-gpt-api/services/auth/business"
	authSQLRepository "lang-gpt-api/services/auth/repository/mysql"
	authUserRPC "lang-gpt-api/services/auth/repository/rpc"
	authAPI "lang-gpt-api/services/auth/transport/api"
	authRPC "lang-gpt-api/services/auth/transport/rpc"
	business2 "lang-gpt-api/services/card/business"
	mysql2 "lang-gpt-api/services/card/repository/mysql"
	api2 "lang-gpt-api/services/card/transport/api"
	"lang-gpt-api/services/gpt/business"
	"lang-gpt-api/services/gpt/repository/mysql"
	"lang-gpt-api/services/gpt/transport/api"
	taskBusiness "lang-gpt-api/services/task/business"
	taskSQLRepository "lang-gpt-api/services/task/repository/mysql"
	taskUserRPC "lang-gpt-api/services/task/repository/rpc"
	taskAPI "lang-gpt-api/services/task/transport/api"
	userBusiness "lang-gpt-api/services/user/business"
	userSQLRepository "lang-gpt-api/services/user/repository/mysql"
	userApi "lang-gpt-api/services/user/transport/api"
	userRPC "lang-gpt-api/services/user/transport/rpc"
)

type TaskService interface {
	CreateTaskHdl() func(*gin.Context)
	GetTaskHdl() func(*gin.Context)
	ListTaskHdl() func(*gin.Context)
	UpdateTaskHdl() func(*gin.Context)
	DeleteTaskHdl() func(*gin.Context)
}

type UserService interface {
	GetUserProfileHdl() func(*gin.Context)
}

type AuthService interface {
	LoginHdl() func(*gin.Context)
	RegisterHdl() func(*gin.Context)
}

type GptService interface {
	CreateMessage() func(c *gin.Context)
	GetListMessage() func(c *gin.Context)
}

type CardService interface {
	CreateCard() gin.HandlerFunc
	GetListCard() gin.HandlerFunc
}

func ComposeUserAPIService(serviceCtx sctx.ServiceContext) UserService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())
	biz := userBusiness.NewBusiness(userRepo)
	userService := userApi.NewAPI(biz)

	return userService
}

func ComposeTaskAPIService(serviceCtx sctx.ServiceContext) TaskService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userClient := taskUserRPC.NewClient(composeUserRPCClient(serviceCtx))
	taskRepo := taskSQLRepository.NewMySQLRepository(db.GetDB())
	biz := taskBusiness.NewBusiness(taskRepo, userClient)
	serviceAPI := taskAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeAuthAPIService(serviceCtx sctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authRepo := authSQLRepository.NewMySQLRepository(db.GetDB())
	hasher := new(common.Hasher)

	userClient := authUserRPC.NewClient(composeUserRPCClient(serviceCtx))
	biz := authBusiness.NewBusiness(authRepo, userClient, jwtComp, hasher)
	serviceAPI := authAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())
	userBiz := userBusiness.NewBusiness(userRepo)
	userService := userRPC.NewService(userBiz)

	return userService
}

func ComposeAuthGRPCService(serviceCtx sctx.ServiceContext) pb.AuthServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authRepo := authSQLRepository.NewMySQLRepository(db.GetDB())
	hasher := new(common.Hasher)

	// In Auth GRPC service, user repository is unnecessary
	biz := authBusiness.NewBusiness(authRepo, nil, jwtComp, hasher)
	authService := authRPC.NewService(biz)

	return authService
}

func ComposeGptService(serviceCtx sctx.ServiceContext) GptService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	gptRepo := mysql.NewMysqlRepoGpt(db)

	biz := business.NewGptBusiness(gptRepo)

	gptService := api.NewAPI(serviceCtx, biz)

	return gptService
}

func ComposeCardService(serviceCtx sctx.ServiceContext) CardService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	cardRepo := mysql2.NewMysqlRepoCard(db)

	biz := business2.NewCardBusiness(cardRepo)

	cardService := api2.NewApi(serviceCtx, biz)

	return cardService
}
