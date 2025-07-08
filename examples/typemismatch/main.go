package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
)

func main() {
	h := map[string]any{"foo": "not-an-int"}
	res := hmap.Get[int](h, "foo")
	if !res.Ok() {
		// res.Err() will be HMapInvalidTypeError
		fmt.Println(res.Err())
	}
}
