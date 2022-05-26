package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Li-Khan/my-forum/config"
	"github.com/Li-Khan/my-forum/repository/postgres"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "config/docker.toml", "path to config.toml file")
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

	router := gin.Default()

	store, err := redis.NewStore(10, "tcp", config.CacheHost+config.CacheAddr, config.CachePassword, []byte("forum"))
	if err != nil {
		log.Println(err)
		return
	}

	router.Use(sessions.Sessions("forum", store))
}
