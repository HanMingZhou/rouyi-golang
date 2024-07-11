package rouyiconf

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func Test_ViperConf(t *testing.T) {
	// 设置配置文件的路径和名称
	viper.AddConfigPath(".")         // "." 表示当前目录
	viper.SetConfigType("yaml")      // 或 "yml"，取决于你的文件扩展名
	viper.SetConfigName("bootstrap") // 文件名，不包含路径和扩展名
	// "." 表示当前目录

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 获取配置值
	proxyEnable := viper.GetBool("go.proxy.enable")

	// 输出结果
	if proxyEnable {
		fmt.Println("Proxy is enabled.")
	} else {
		fmt.Println("Proxy is disabled.")
	}
}
