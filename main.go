package main

import (
	"github.com/reiver/xim-id/arg"

	"github.com/reiver/go-xim"

	"fmt"
	"os"
)

func main() {

	var command string
	{
		if 1 <= len(arg.Values) {
			command = arg.Values[0]
		}
	}

	switch command {
	case "generate", "":
		generate()
	default:
		fmt.Fprintf(os.Stderr, "error: unknown command — %q\n", command)
		os.Exit(1)
	}
}


func generate() {
	var id xim.ID = xim.Generate()

	switch {
	case arg.NoNewline:
		fmt.Print(id)
	default:
		fmt.Println(id)
	}
}
