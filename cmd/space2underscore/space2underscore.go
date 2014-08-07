package main

import (
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	p := os.Args
	string := strings.Join(p[1:], "_")
	if err := clipboard.WriteAll(string); err != nil {
		panic(err)
	}
}
