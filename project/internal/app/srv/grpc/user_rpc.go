package grpc

import (
	"context"
	"project/api/v1/proto"
)

// UserService 定义 UserService 并实现约定的接口
type UserService struct {
	proto.UnimplementedUserServer `wire:"-"`
}

func (h UserService) GetUserById(ctx context.Context, id *proto.IdRequest) (*proto.UserInfoResponse, error) {
	resp := new(proto.UserInfoResponse)
	resp.Id = 1
	return resp, nil
}
