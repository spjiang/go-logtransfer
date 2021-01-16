package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logtransfer/conf"
	"logtransfer/es"
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
	fmt.Println("init conf success.")

	// 1、初始化es
	err = es.Init(cfg.ESConfig.Address, cfg.ESConfig.ChanSize, cfg.ESConfig.ChanWorker)
	if err != nil {
		fmt.Printf("Init es failed, err:%v\n", err)
		return
	}
	fmt.Println("init es success.")

	// 2、初始化 kafka
	// 2.1 连接kafka,创建分区的消费者
	// 2.2 每个分区的消费者分布取出数据，通过SendToES()将数据发往ES
	err = kafka.Init([]string{cfg.KafkaConfig.Address}, cfg.KafkaConfig.Topic)
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")
	select {}
}
