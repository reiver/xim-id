package main

import (
	"github.com/reiver/xim-id/arg"
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
	case "decompile":
		decompile()
	default:
		badrequest("error: unknown command — %q", command)
		return
	}
}
