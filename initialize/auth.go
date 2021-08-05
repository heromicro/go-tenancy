package initialize

import (
	"os"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

func Auth() {
	err := multi.InitDriver(&multi.Config{
		DriverType:      g.TENANCY_CONFIG.System.CacheType,
		UniversalClient: g.TENANCY_CACHE})
	if err != nil {
		g.TENANCY_LOG.Error("初始化缓存驱动:", zap.Any("err", err))
	}

	if multi.AuthDriver == nil {
		g.TENANCY_LOG.Error("初始化缓存驱动失败")
		os.Exit(0)
	}
}
