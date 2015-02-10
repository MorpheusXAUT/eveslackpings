package main

import (
	"log"
	"os"
	"runtime"

	"github.com/morpheusxaut/eveslackpings/misc"
	"github.com/morpheusxaut/eveslackpings/web"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config, err := misc.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: [%v]", err)
		os.Exit(2)
	}

	misc.SetupLogger(config.DebugLevel)

	controller := web.SetupController(config)

	controller.HandleRequests()
}
