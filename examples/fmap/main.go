package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
	"github.com/zodimo/go-hmap-helper/hmap/result"
)

func main() {
	h := map[string]any{"foo": 42}
	f := func(x int) result.HMapResult[int] { return result.NewMapValidResult(x * 10) }
	res := hmap.FMap[int, int](h, "foo", f)
	if res.Ok() {
		v, _ := res.Value()
		fmt.Println(v) // 420 if h["foo"] == 42
	}
}
