package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka日志模块
type logData struct {
	topic string
	data  string
}

var (
	consumer sarama.Consumer // 声明一个全局的一个kafka生产者客户端
)

// Init 初始化client
func Init(addrs []string, topic string) (err error) {
	// 连接kafka
	consumer, err = sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer failed,err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of Partitions failed,err:%v\n", err)
		return
	}
	fmt.Printf("get Partitions :%v\n", partitionList)

	// 便利所有的分区
	for partition := range partitionList {
		// 对应每个分区创建一个消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d, offset:%d, key:%v,Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				// 直接发送给ES
			}
		}(pc)
	}

	return
}
