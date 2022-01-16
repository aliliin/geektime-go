//go:build wireinject
// +build wireinject

package main

import (
	"github.com/quexer/utee"
	"log"
	"project/internal/app/srv"
)

func main() {
	app, err := srv.RunSrv()
	utee.Chk(err)
	err = app.Init().Run()
	if err != nil {
		log.Fatalln(err)
	}
}
