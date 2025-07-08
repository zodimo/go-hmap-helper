package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
)

func main() {
	h := map[string]any{"foo": 42}
	res := hmap.Get[int](h, "foo")
	if res.Ok() {
		v, _ := res.Value()
		fmt.Println(v) // 42
	} else {
		fmt.Println("Error:", res.Err())
	}
}
