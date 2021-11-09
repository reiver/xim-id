package main

import (
	"github.com/reiver/xim-id/arg"
	"github.com/reiver/xim-id/lib/opt/int32"

	"github.com/reiver/go-xim"

	"fmt"
)

func generate() {

	var id xim.ID
	{
		switch {
		case optint32.Nothing() != arg.UnixTime:
			unixtimestamp, successful := arg.UnixTime.Return()
			if !successful {
				internalerror("was not able to get unix-time")
				return
			}

			id = xim.GenerateForUnixTime(unixtimestamp)
		default:
			id = xim.Generate()
		}
	}

	switch {
	case arg.NoNewline:
		fmt.Print(id)
	default:
		fmt.Println(id)
	}
}
