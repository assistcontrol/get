package main

import (
	"log"
	"os"

	"github.com/assistcontrol/get/config"
	"github.com/assistcontrol/get/fetch"
	"github.com/assistcontrol/get/output"
)

func main() {
	conf, err := config.NewConfig(os.Args)
	if err != nil {
		log.Printf("Error: %v\n%s", err, help())
		os.Exit(1)
	}

	contents, err := fetch.Fetch(conf)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if conf.ShouldSave() {
		output.Get(contents)
	} else {
		output.Show(contents)
	}
}

func help() string {
	return "Usage: get <url>"
}
