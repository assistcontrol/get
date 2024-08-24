package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"golang.org/x/term"
)

const (
	// Default values for Chroma
	ChromaDefaultFT = "html"
	ChromaFormat    = "terminal16m"
	ChromaStyle     = "catppuccin-mocha"
)

func show(bytes []byte) {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println(colorize(bytes))
		return
	}

	fmt.Println(string(bytes))
}

func colorize(raw []byte) string {
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
		return contents
	}

	var b bytes.Buffer
	err = formatter.Format(&b, style, iterator)
	if err != nil {
		return contents
	}

	return b.String()
}
