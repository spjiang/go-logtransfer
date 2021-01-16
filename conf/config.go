package conf

type LogTransfer struct {
	KafkaConfig `ini:"kafka"`
	ESConfig    `ini:"es"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESConfig struct {
	Address    string `ini:"address"`
	ChanSize   int    `ini:"chanSize"`
	ChanWorker int    `ini:"chanWorker"`
}
