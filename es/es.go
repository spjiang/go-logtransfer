package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"strings"
	"time"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

// 初始化ES，准备接收kafka数据
// Init ...
func Init(addr string, chanSize, chanWorker int) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		fmt.Println("++++++")
		return err
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogData, chanSize)
	// 启动多个goroutine去消费通道数据
	for i := 0; i < chanWorker; i++ {
		go SendToES()
	}
	return
}

// SendToCh 日志数据写入通道
func SendToESChan(msg *LogData) {
	ch <- msg
}

// SendToES 发送数据到ES
func SendToES() {
	for {
		select {
		case msg := <-ch:
			_, err := client.Index().Index(msg.Topic).Type("XXXX").BodyJson(msg).Do(context.Background())
			if err != nil {
				continue
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
