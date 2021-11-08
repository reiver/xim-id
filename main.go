package main

import (
	"github.com/reiver/xim-id/arg"

	"github.com/reiver/go-xim"

	"fmt"
)

func main() {

	var id xim.ID = xim.Generate()

	switch {
	case arg.NoNewline:
		fmt.Print(id)
	default:
		fmt.Println(id)
	}
}
