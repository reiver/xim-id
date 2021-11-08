package arg

import (
	"flag"
)

var (
	NoNewline bool
)


func init() {
	flag.BoolVar(&NoNewline, "n", false, "do not output the trailing newline")

	flag.Parse()
}
