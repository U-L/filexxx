package global

import (
	"filexxx/config"
)

var (
	MyConfig *config.MyConfig = &config.MyConfig{
		Host: "127.0.0.1",
		Port: 8090,
	}


	Address *config.Address = &config.Address{}

	MyIP string
	MyIPTmp string
)
