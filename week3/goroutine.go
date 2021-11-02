package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

/*
问题描述：
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
根据描述信息，可以简单汇总成3块内容：

1.实现 HTTP server 的启动和关闭
2.监听 linux signal 信号，使用chan实现对 linux signal中断的注册和处理 按 ctrl+c 之类退出程序
3.errgroup 实现多个 goroutine 的级联退出，通过 errgroup+context 的形式，对 1、2 中的 goroutine 进行级联注销
*/
func main() {
	group, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Go")
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8889",
	}

	// 利用无缓冲chan 模拟单个服务错误退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{} // 阻塞
	})

	// -- 测试 http server 的启动和退出 --

	// g1 启动http server服务
	// g1 退出后, context 将不再阻塞，g2, g3 都会随之退出
	group.Go(func() error {
		return server.ListenAndServe()
	})

	// g2
	// g2 退出时，调用了 shutdown，g1 也会退出
	group.Go(func() error {
		select {
		case <-serverOut:
			fmt.Println("server closed") // 退出会触发 g.cancel, ctx.done 会收到信号
		case <-ctx.Done():
			fmt.Println("errgroup exit")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx)
	})

	// g3 linux signal 信号的注册和处理
	// g3 捕获到 os 退出信号将会退出
	// g3 退出后, context 将不再阻塞，g2 会随之退出
	// g2 退出时，调用了 shutdown，g1 会退出
	group.Go(func() error {
		quit := make(chan os.Signal, 1)
		// sigint 用户ctrl+c, sigterm程序退出
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os exit: %v", sig)
		}
	})

	// 然后 main 函数中的 g.Wait() 退出，所有协程都会退出
	err := group.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
}

func httpServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	if err := http.ListenAndServe("0.0.0.0:8089", mux); err != nil {
		return err
	}
	return nil
}

func debugServer() error {
	if err := http.ListenAndServe("127.0.0.1:8088", nil); err != nil {
		log.Fatal(err)
	}
	return nil
}
