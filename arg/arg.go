package arg

import (
	"github.com/reiver/xim-id/lib/opt/int32"

	"flag"
)

var (
	Values []string
)

var (
	NoNewline bool
	UnixTime optint32.OptInt32
)

func init() {
	flag.BoolVar(&NoNewline, "n", false, "do not output the trailing newline after xim-id")
	flag.Var(&UnixTime, "unixtime", "unix-time to use with generation of xim-id")

	flag.Parse()

	Values = flag.Args()
}
