package arg

import (
	"flag"
)

var (
	Values []string
)

var (
	NoNewline bool
)

func init() {
	flag.BoolVar(&NoNewline, "n", false, "do not output the trailing newline")

	flag.Parse()

	Values = flag.Args()
}
