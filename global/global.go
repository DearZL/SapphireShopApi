package global

import (
	"SapphireShop/SapphireShop_api/common/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Config *viper.Viper
	Logger *zap.Logger
)

func InitConfig() {
	Config = config.InitConfig()
}

func InitLogConfig() {
	if Config.Get("app.mode") == "debug" {
		Logger = config.ZapConfig(true)
	}

}
