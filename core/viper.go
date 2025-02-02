package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用 config 的默认值, config 的路径为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用 TENANCY_CONFIG 环境变量, config 的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config 的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用 func Viper() 传递的值,config 的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&g.TENANCY_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&g.TENANCY_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
