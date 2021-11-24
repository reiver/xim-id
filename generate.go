package main

import (
	"github.com/reiver/xim-id/arg"
	"github.com/reiver/xim-id/lib/opt/int32"

	"github.com/reiver/go-xim"

	"fmt"
	"sort"
)

func generate() {

	var ids []xim.ID
	{
		switch {
		case optint32.Nothing() != arg.UnixTime:
			unixtimestamp, successful := arg.UnixTime.Return()
			if !successful {
				internalerror("was not able to get unix-time")
				return
			}

			var id xim.ID = xim.GenerateForUnixTime(unixtimestamp)

			ids = append(ids, id)
		default:
			for i:=uint64(0); i<arg.Count; i++ {
				var id xim.ID = xim.Generate()
				ids = append(ids, id)
			}
		}
	}

	{
		var fn func(i int, j int)bool = func(i int, j int) bool {
			return ids[i].String() < ids[j].String()
		}

		sort.Slice(
			ids,
			fn,
		)
	}

	switch {
	case arg.NoNewline:
		for i, id := range ids {
			if 0 < i {
				fmt.Print(" ")
			}
			fmt.Print(id)
		}
	default:
		for _, id := range ids {
			fmt.Println(id)
		}
	}
}
