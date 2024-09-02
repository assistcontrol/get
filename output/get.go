package output

import (
	"log"
	"os"

	"github.com/assistcontrol/get/context"
)

func Get(c *context.Ctx) {
	flags := os.O_WRONLY | os.O_CREATE
	switch c.Force {
	case true:
		flags |= os.O_TRUNC
	case false:
		flags |= os.O_EXCL
	}

	f, err := os.OpenFile(c.Destination, flags, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.Write(c.Body); err != nil {
		log.Fatal(err)
	}
}
