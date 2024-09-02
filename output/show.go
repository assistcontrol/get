package output

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/assistcontrol/get/context"
	"golang.org/x/term"
)

const (
	// Default values for Chroma
	chromaDefaultFT = "html"
	chromaFormat    = "terminal16m"
	chromaStyle     = "catppuccin-mocha"
)

// show either colorizes and prints the body to the terminal
// (if it's a terminal), or prints it as-is.
func show(c *context.Ctx) error {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		c.Body = colorize(c.Body)
	}

	fmt.Println(string(c.Body))

	return nil
}

// colorize colorizes the raw byte slice. If anything goes wrong,
// it just returns the raw input.
func colorize(raw []byte) []byte {
	contents := string(raw)

	lexer := lexers.Analyse(contents)
	if lexer == nil {
		lexer = lexers.Get(chromaDefaultFT)
	}
	lexer = chroma.Coalesce(lexer)

	style := styles.Get(chromaStyle)
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get(chromaFormat)
	if formatter == nil {
		formatter = formatters.Fallback
	}

	iterator, err := lexer.Tokenise(nil, contents)
	if err != nil {
		return raw
	}

	var b bytes.Buffer
	err = formatter.Format(&b, style, iterator)
	if err != nil {
		return raw
	}

	return b.Bytes()
}
