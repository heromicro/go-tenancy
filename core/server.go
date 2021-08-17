package core

import (
	"embed"
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer(fs embed.FS) {
	Router := initialize.App(fs)
	// Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", g.TENANCY_CONFIG.System.Addr)
	// In order to ensure that the text order output can be
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	g.TENANCY_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf("默认监听地址:http://127.0.0.1%s\n", address)
	g.TENANCY_LOG.Error(s.ListenAndServe().Error())
}
