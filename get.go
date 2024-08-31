package main

import "os"

func get(contents []byte) {
	os.Stdout.Write(contents)
}
