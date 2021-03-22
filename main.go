package main

import (
	"log"

	"github.com/manojown/restApi/app"
	"github.com/manojown/restApi/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.NewConfig()
	app.ConfigAndRun(config)

}
