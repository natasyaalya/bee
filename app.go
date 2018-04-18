package main

import (
	"flag"
	"log"

	"github.com/asepnur/iskandar/src/util/conn"
	"github.com/asepnur/iskandar/src/util/env"
	"github.com/asepnur/iskandar/src/util/jsonconfig"
	"github.com/asepnur/iskandar/src/webserver"
)

var (
	topic  = "180204"
	chanel = "iskandar"
)

type configuration struct {
	Database  conn.DatabaseConfig `json:"database"`
	Redis     conn.RedisConfig    `json:"redis"`
	Webserver webserver.Config    `json:"webserver"`
}

func main() {
	flag.Parse()

	// load config
	cfgenv := env.Get()
	config := &configuration{}
	isLoaded := jsonconfig.Load(&config, "/etc", cfgenv) || jsonconfig.Load(&config, "./files/etc", cfgenv)
	if !isLoaded {
		log.Fatal("Failed to load configuration")
	}
	// initialize instance
	conn.InitRedis(config.Redis)
	conn.Consume(topic, chanel)
	conn.InitDB(config.Database)
	webserver.Start(config.Webserver)
}
