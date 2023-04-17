package rpc

import (
	"context"
	"github.com/pkg/errors"
	"lang-gpt-api/proto/pb"
)

type rpcClient struct {
	client pb.UserServiceClient
}

func NewClient(client pb.UserServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error) {
	resp, err := c.client.CreateUser(ctx, &pb.CreateUserReq{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	})

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(resp.Id), nil
}
