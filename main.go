package main

import (
	"flag"
	"log"
	"os"

	"github.com/assistcontrol/get/context"
	"github.com/assistcontrol/get/fetch"
	"github.com/assistcontrol/get/output"
)

func main() {
	// Parse command line arguments
	ctx, err := context.New()
	if err != nil {
		usage()
	}

	// Fetch the requested resource
	err = fetch.Fetch(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Display or save it
	if ctx.Saving {
		output.Get(ctx)
	} else {
		output.Show(ctx)
	}
}

func usage() {
	flag.Usage()
	os.Exit(1)
}
