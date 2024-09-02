package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/assistcontrol/get/context"
	"github.com/assistcontrol/get/fetch"
	"github.com/assistcontrol/get/output"
)

func main() {
	// Parse command line arguments
	ctx, err := context.New()
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	// Fetch the requested resource
	if err = fetch.Fetch(ctx); err != nil {
		die(err)
	}

	// Display or save it
	if err = output.Output(ctx); err != nil {
		die(err)
	}
}

func die(e error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", e)
	os.Exit(1)
}
