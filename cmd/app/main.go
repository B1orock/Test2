package main

import (
	"Varian_v2/config"
	"log"
)

func main() {

	cfg, err := config.GetENVConfig()
	if err != nil {
		log.Fatal("error load config", err)
	}
	log.Println(cfg)
}
