package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"golang.org/x/term"
)

var (
	URLPattern = regexp.MustCompile(`^https?://`)
)

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
		return
	}
	url := args[1]

	contents, err := os.ReadFile(url)
	if err != nil {
		contents, err = get(url)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	show(contents)
}

func help() {
	fmt.Println("Usage: get <url>")
}

func get(url string) ([]byte, error) {
	if URLPattern.MatchString(url) {
		contents, err := fetch(url)
		if err != nil {
			return nil, err
		}
		return contents, nil
	}

	contents, err := fetch("https://" + url)
	if err != nil {
		contents, err = fetch("http://" + url)
	}
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

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
		lexer = lexers.Get("html")
	}
	fmt.Printf("Lexer: %s\n", lexer.Config().Name)
	lexer = chroma.Coalesce(lexer)

	style := styles.Get("catppuccin-mocha")
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get("terminal16m")
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
