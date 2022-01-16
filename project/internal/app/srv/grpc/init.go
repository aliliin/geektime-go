package grpc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"project/api/v1/proto"
)

// Initializer grpc 初始化器
type Initializer struct {
	Config      *viper.Viper
	UserService *UserService
}

func (p *Initializer) Name() string {
	return "grpc_initializer"
}

func (p *Initializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

// Initialize 进行micro grpc 注册
func (p *Initializer) Initialize(ctx context.Context) error {
	go p.startRpc()
	return nil
}

func (p *Initializer) startRpc() error {
	listen, err := net.Listen("tcp", p.Config.GetString("grpc.port"))

	fmt.Println(p.Config.GetString("grpc.port"))
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册UserService
	proto.RegisterUserServer(s, p.UserService)

	s.Serve(listen)
	return nil
}
