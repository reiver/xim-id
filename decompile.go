package main

import (
	"github.com/reiver/xim-id/arg"

	"github.com/reiver/go-xim"

	"fmt"
	"time"
)

func decompile() {
	if len(arg.Values) < 2 {
		badrequest("missing xim-id")
		return
	}

	var ximnotation string = arg.Values[1]

	var id xim.ID
	{
		err := id.UnmarshalText([]byte(ximnotation))
		if nil != err {
			badrequest("invalid xim-id — %q", ximnotation)
			return
		}
	}

	{
		var unixtimestamp int64
		{
			var successful bool

			unixtimestamp, successful = id.UnixTime()
			if !successful {
				internalerror("internal software error: was not able to decompile unix timestamp")
				return
			}
		}

		var t time.Time
		{
			t = time.Unix(unixtimestamp, 0)
		}

		var chaos uint64
		{
			var successful bool

			chaos, successful = id.Chaos()
			if !successful {
				internalerror("internal software error: was not able to decompile chaos")
				return
			}
		}

		fmt.Printf("time:\t%v (unix-time = %v)\n", t, unixtimestamp)
		fmt.Printf("chaos:\t%#020b\n", chaos)
	}
}
