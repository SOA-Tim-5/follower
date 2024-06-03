package config

type Config struct {
	Address                   string
	NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
	CreateOrderCommandSubject string
	CreateOrderReplySubject   string
}

func GetConfig() Config {
	return Config{
		Address:                   ":44333",
		NatsHost:                  "localhost",
		NatsPort:                  "4222",
		NatsUser:                  "ruser",
		NatsPass:                  "T0pS3cr3t",
		CreateOrderCommandSubject: "order.create.command",
		CreateOrderReplySubject:   "order.create.reply",
	}
}
