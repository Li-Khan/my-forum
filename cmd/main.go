package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Li-Khan/my-forum/config"
	"github.com/Li-Khan/my-forum/repository/postgres"
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

	db, err := postgres.NewPostgresRepository(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
}
