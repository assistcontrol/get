package output

import (
	"fmt"
	"os"

	"github.com/assistcontrol/get/body"
)

func Get(b *body.Body) {
	os.Stdout.Write(b.Body)
	fmt.Println(b.Filename)
}
