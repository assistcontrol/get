package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/assistcontrol/get/config"
	"github.com/assistcontrol/get/fetch"
	"github.com/assistcontrol/get/output"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		showHelp()
		os.Exit(1)
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

func showHelp() {
	fmt.Fprintf(os.Stderr, "Usage: get [flags] <URL>\n")
	flag.PrintDefaults()
}
