package initialize

import (
	"os"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

func Auth() {
	g.TENANCY_LOG.Info("初始化认证驱动", zap.String("使用缓存类型", g.TENANCY_CONFIG.System.CacheType))
	if g.TENANCY_CONFIG.System.CacheType == "" {
		return
	}
	if g.TENANCY_CONFIG.System.CacheType == "redis" {
		g.TENANCY_CACHE = cache.Cache() // redis缓存
	}
	err := multi.InitDriver(&multi.Config{
		DriverType:      g.TENANCY_CONFIG.System.CacheType,
		UniversalClient: g.TENANCY_CACHE})
	if err != nil {
		g.TENANCY_LOG.Error("初始化缓存驱动:", zap.Any("err", err))
	}

	if g.TENANCY_CONFIG.System.CacheType != "local" && multi.AuthDriver == nil {
		g.TENANCY_LOG.Error("初始化缓存驱动失败")
		os.Exit(0)
	}
}
