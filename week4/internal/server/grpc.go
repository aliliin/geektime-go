package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "week4/api/v1/proto"
	biz "week4/internal/biz"
	"week4/internal/data"
	model "week4/internal/model"
)

type GrpcServer struct {
	pb.UnimplementedUserServer
	userService *biz.UserService
	svr         *grpc.Server
}

func NewGrpcServer(userService *biz.UserService) *GrpcServer {
	return &GrpcServer{
		userService: userService,
	}
}

func (server *GrpcServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.Response, error) {
	if req.UserName == "" {
		return &pb.Response{
			Code:    400,
			Message: "用户名不能为空",
		}, nil
	}

	if req.Password == "" {
		return &pb.Response{
			Code:    400,
			Message: "密码必填",
		}, nil
	}

	if req.Email == "" {
		return &pb.Response{
			Code:    400,
			Message: "email can not be empty",
		}, nil
	}

	user := model.User{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	}
	result, err := server.userService.Register(ctx, user)
	if err != nil {
		if errors.Is(err, data.ErrAlreadyExist) {
			return &pb.Response{
				Code:    400,
				Message: fmt.Sprintf("user with user name %s is already exist", req.UserName),
			}, nil
		} else {
			fmt.Printf("GrpcServer error, error while create user, %s", err)
			return &pb.Response{
				Code:    500,
				Message: "internal server error",
			}, nil
		}
	}
	if result {
		return &pb.Response{
			Code:    200,
			Message: "success",
			Data:    true,
		}, nil
	} else {
		return &pb.Response{
			Code:    200,
			Message: "failed",
			Data:    false,
		}, nil
	}
}

func (server *GrpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	if req.UserName == "" {
		return &pb.Response{
			Code:    400,
			Message: "user name can not be empty",
		}, nil
	}

	if req.Password == "" {
		return &pb.Response{
			Code:    400,
			Message: "password can not be empty",
		}, nil
	}
	result, err := server.userService.Login(ctx, req.UserName, req.Password)
	if err != nil {
		if errors.Is(err, data.ErrNotExist) {
			return &pb.Response{
				Code:    200,
				Message: "could find user",
			}, nil
		} else {
			fmt.Printf("GrpcServer error, error while login user, %s", err)
			return &pb.Response{
				Code:    500,
				Message: "internal server error",
			}, nil
		}
	}
	if result {
		return &pb.Response{
			Code:    200,
			Message: "success",
			Data:    true,
		}, nil
	} else {
		return &pb.Response{
			Code:    200,
			Message: "password error",
		}, nil
	}
}

func (server *GrpcServer) ListenAndServe(ctx context.Context, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server.svr = grpcServer
	pb.RegisterUserServer(grpcServer, server)
	fmt.Printf("start grpc server at %s\n", address)
	return grpcServer.Serve(lis)
}

func (server *GrpcServer) Shutdown(ctx context.Context) error {
	server.svr.GracefulStop()
	return nil
}
