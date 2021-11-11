package server

import (
	"context"
	"fmt"

	"week4/config"

	"golang.org/x/sync/errgroup"
)

type Server interface {
	ListenAndServe(ctx context.Context, address string) error
	Shutdown(ctx context.Context) error
}

type UserServer struct {
	grpcServer *GrpcServer
	httpServer *HttpServer
	config     *config.Config
	ctx        context.Context
	cancelF    context.CancelFunc
}

func NewUserServer(grpc *GrpcServer, http *HttpServer, config *config.Config) *UserServer {
	return &UserServer{
		grpcServer: grpc,
		httpServer: http,
		config:     config,
	}
}

func (s *UserServer) Start(ctx context.Context) error {
	newContext, cancelF := context.WithCancel(ctx)
	s.ctx = newContext
	s.cancelF = cancelF
	group, cctx := errgroup.WithContext(s.ctx)
	group.Go(func() error {
		go func() {
			<-ctx.Done()
			fmt.Printf("context done receive, exiting http stop\n")
			s.httpServer.Shutdown(cctx)
		}()
		return s.httpServer.ListenAndServe(cctx, s.config.HttpAddress)
	})
	group.Go(func() error {
		go func() {
			<-ctx.Done()
			fmt.Printf("context done receive, exiting grpc stop;\n")
			_ = s.grpcServer.Shutdown(cctx)

		}()
		return s.grpcServer.ListenAndServe(cctx, s.config.GrpcAddress)
	})
	return group.Wait()
}

func (s *UserServer) Shutdown() {
	s.cancelF()
}
