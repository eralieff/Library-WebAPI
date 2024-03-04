package main

import (
	"Library_WebAPI/internal/app"
	"Library_WebAPI/pkg/config"
	"log"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println("Cannot load configs", err)
		return
	}

	err = app.Start(conf)
	if err != nil {
		log.Println("Cannot Start server")
		return
	}
}
