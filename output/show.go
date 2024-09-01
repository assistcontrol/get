package output

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/assistcontrol/get/body"
	"golang.org/x/term"
)

const (
	// Default values for Chroma
	ChromaDefaultFT = "html"
	ChromaFormat    = "terminal16m"
	ChromaStyle     = "catppuccin-mocha"
)

func Show(b *body.Body) {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		b.Body = colorize(b.Body)
	}

	fmt.Println(string(b.Body))
}

func colorize(raw []byte) []byte {
	contents := string(raw)

	lexer := lexers.Analyse(contents)
	if lexer == nil {
		lexer = lexers.Get(ChromaDefaultFT)
	}
	lexer = chroma.Coalesce(lexer)

	style := styles.Get(ChromaStyle)
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get(ChromaFormat)
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
