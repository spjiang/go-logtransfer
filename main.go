package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logtransfer/conf"
	"logtransfer/kafka"
)

// log transfer
// 将日志数据从kafka写入es
var cfg = new(conf.LogTransfer)

func main() {
	// 0、获取配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("Init config failed, err:%v\n", err)
		return
	}
	// 1、初始化 kafka连接
	err = kafka.Init([]string{cfg.KafkaConfig.Address}, cfg.KafkaConfig.Topic)
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	// 2、写入es
}
