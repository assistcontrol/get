package output

import (
	"log"
	"os"

	"github.com/assistcontrol/get/body"
	"github.com/assistcontrol/get/config"
)

func Get(b *body.Body, c *config.Config) {
	path := c.Filename
	if path == "" {
		path = b.Filename
	}

	flags := os.O_WRONLY | os.O_CREATE
	switch c.Force {
	case true:
		flags |= os.O_TRUNC
	case false:
		flags |= os.O_EXCL
	}

	f, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.Write(b.Body); err != nil {
		log.Fatal(err)
	}
}
