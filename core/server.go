package core

import (
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.App()
	address := fmt.Sprintf(":%d", g.TENANCY_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	g.TENANCY_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf("默认监听地址:http://127.0.0.1%s\n", address)
	g.TENANCY_LOG.Error(s.ListenAndServe().Error())
}
