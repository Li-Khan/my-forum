package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Li-Khan/my-forum/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "config/server.toml", "path to config.toml file")
}

func main() {
	flag.Parse()

	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println(err)
		return
	}
}
