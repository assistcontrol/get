// Package output provides functions that write to a destination file,
// or colorize and write to the terminal.
package output

import (
	"os"

	"github.com/assistcontrol/get/context"
)

// Output looks at the context and decides what to do with the body.
// If the save flag is set, it writes to the destination file.
// Otherwise, it colorizes the body and writes it to the terminal.
// It returns an error if something goes wrong.
func Output(c *context.Ctx) error {
	if c.Save {
		return save(c)
	}

	return show(c)
}

// save writes the body to the destination file. It will overwrite
// existing files if the force flag is set. If the file already exists
// and the force flag is not set, it will not overwrite the file.
// It returns an error if the file cannot be written.
func save(c *context.Ctx) error {
	flags := os.O_WRONLY | os.O_CREATE
	switch c.Force {
	case true:
		flags |= os.O_TRUNC
	case false:
		flags |= os.O_EXCL
	}

	f, err := os.OpenFile(c.Destination, flags, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic("Error closing output file: " + err.Error())
		}
	}()

	_, err = f.Write(c.Body)

	return err
}
