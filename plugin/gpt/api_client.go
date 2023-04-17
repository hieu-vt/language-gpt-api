package gpt

import (
	"context"
	"flag"
	"fmt"
	"github.com/sashabaranov/go-openai"
	sctx "github.com/viettranx/service-context"
)

type gpt struct {
	id           string
	client       *openai.Client
	logger       sctx.Logger
	apiToken     string
	organization string
}

func NewGptClient(id string) *gpt {
	return &gpt{id: id}
}

func (g *gpt) ID() string {
	return g.id
}

func (g *gpt) InitFlags() {
	flag.StringVar(&g.apiToken, "GPT_API_TOKEN", "", "Api token of chatgpt")
	flag.StringVar(&g.organization, "GPT_organization", "", "Api token of chatgpt")
}

func (g *gpt) Activate(context sctx.ServiceContext) error {
	g.logger = sctx.GlobalLogger().GetLogger(g.id)
	g.client = g.configGptAPI()

	return nil
}

func (g *gpt) Stop() error {
	return nil
}

func (g *gpt) configGptAPI() *openai.Client {
	if g.apiToken == "" {
		return nil
	}
	
	client := openai.NewClient(g.apiToken)

	g.logger.Infoln("Started new client gpt")

	return client
}

func (g *gpt) RequestGptAPI(ctx context.Context, message string) (error, *ResponseType) {
	resp, err := g.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	return nil, &ResponseType{Message: resp.Choices[0].Message.Content}
}
