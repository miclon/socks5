package main

import (
	"fmt"
	"log"

	"github.com/armon/go-socks5"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件变更:", e.Name)
		// 在配置文件变更时手动刷新配置
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("刷新配置失败:", err)
		} else {
			fmt.Println("配置已刷新")
		}
	})
	config := &socks5.Config{}
	server, err := socks5.New(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Socks5服务已启动")
	if err := server.ListenAndServe("tcp", ":"+viper.GetString("port")); err != nil {
		log.Fatal(err)
	}
}
