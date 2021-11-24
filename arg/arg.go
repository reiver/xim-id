package arg

import (
	"github.com/reiver/xim-id/lib/opt/int32"
	"github.com/reiver/xim-id/lib/opt/uint64"

	"flag"
)

var (
	Values []string
)

var (
	Count     uint64
	NoNewline bool
	UnixTime  optint32.OptInt32
)

var (
	count optuint64.OptUint64
)

func init() {

	flag.Var(&count, "count", "the number of xim-ids to generate")
	flag.BoolVar(&NoNewline, "n", false, "do not output the trailing newline after xim-id")
	flag.Var(&UnixTime, "unixtime", "unix-time to use with generation of xim-id")

	flag.Parse()

	Values = flag.Args()
	Count = count.ForcedReturn(1)
}
