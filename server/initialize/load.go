package initialize

import (
	"fmt"
	"nocake/global"

	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件
func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Fatal err resources file :%s", err.Error())
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Fatal err resources file :%s", err.Error())
	}
}
