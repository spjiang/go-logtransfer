package conf

type LogTransfer struct {
	KafkaConfig `ini:"kafka"`
	ESConfig    `ini:"etcd"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESConfig struct {
	Address string `ini:"address"`
}
