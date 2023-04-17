package gpt

import "context"

type ResponseType struct {
	Message string `json:"message"`
}

type GptClient interface {
	RequestGptAPI(ctx context.Context, message string) (error, *ResponseType)
}
