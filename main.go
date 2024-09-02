package main

import (
	"flag"
	"log"
	"os"

	"github.com/assistcontrol/get/config"
	"github.com/assistcontrol/get/fetch"
	"github.com/assistcontrol/get/output"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		usage()
	}

	contents, err := fetch.Fetch(conf)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if conf.Saving {
		output.Get(contents, conf)
	} else {
		output.Show(contents, conf)
	}
}

func usage() {
	flag.Usage()
	os.Exit(1)
}
