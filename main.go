package main

import (
	"github.com/reiver/go-xim"

	"fmt"
)

func main() {

	var id xim.ID = xim.Generate()

	fmt.Println(id)
}
