package initialize

import (
	"context"
	"net"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

func Auth() {
	universalOptions := &redis.UniversalOptions{
		Addrs:       strings.Split(g.TENANCY_CONFIG.Redis.Addr, ","),
		Password:    g.TENANCY_CONFIG.Redis.Password,
		PoolSize:    10,
		IdleTimeout: 300 * time.Second,
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := net.Dial(network, addr)
			if err == nil {
				go func() {
					time.Sleep(5 * time.Second)
					conn.Close()
				}()
			}
			return conn, err
		},
	}
	err := multi.InitDriver(&multi.Config{
		DriverType:       g.TENANCY_CONFIG.System.CacheType,
		UniversalOptions: universalOptions})
	if err != nil {
		g.TENANCY_LOG.Error("初始化缓存驱动:", zap.Any("err", err))
	}

	if multi.AuthDriver == nil {
		g.TENANCY_LOG.Error("初始化缓存驱动失败")
		os.Exit(0)
	}
}
