package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
)

func main() {
	h := map[string]any{"foo": 42}
	res := hmap.Get[int](h, "bar")
	if !res.Ok() {
		// res.Err() will be HMapKeyNotFoundError
		fmt.Println(res.Err())
	}
}
