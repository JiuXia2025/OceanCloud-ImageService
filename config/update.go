package config

import (
	"fmt"
	"log"
)

func UpdateCheck() {
	//假设云端下发的最新版本
	var newversion string = "1.0"
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	//fmt.Printf("当前程序版本: %s\n", cfg.Version)
	if cfg.Version >= newversion {
		fmt.Println("[Update/INFO]", cfg.AppName, "当前版本为最新版本", cfg.Version, ",无需更新")
	} else if cfg.Version < newversion {
		fmt.Println("[Update/INFO] 最新版已出，建议尽快更新服务端")
	}
}
func UpdateAPP() {
	//此处执行下发热更新版本
}
