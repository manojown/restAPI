package main

import (
	"log"

	"github.com/restApi/app"
	"github.com/restApi/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.NewConfig()
	app.ConfigAndRun(config)

}
