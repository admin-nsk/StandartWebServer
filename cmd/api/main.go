package main

import (
	"flag"
	"github.com/admin-nsk/StandartWebServer/internal/app/api"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath   string = "configs/api.toml"
	configFormat string = "toml"
)

func init() {
	flag.StringVar(&configPath, "config", "configs/api.toml",
		"path to config file in .toml or .env format")
	flag.StringVar(&configFormat, "format", "toml", "format of config file")
}

func main() {
	flag.Parse()
	log.Println("Start server..")
	//Server instance initialization
	config := api.NewConfig()
	var err error
	if configFormat == "toml" {
		_, err = toml.DecodeFile(configPath, config)

	} else {
		api.GetEnvConfig(configPath, config)
	}
	if err != nil {
		log.Println("Can not read config file. Using default values: ", err)
	}
	server := api.New(config)

	//api server start
	log.Fatal(server.Start())
}
