package hmap

import (
	"reflect"

	"github.com/zodimo/go-hmap-helper/hmap/result"
)

func FMap[IN, OUT any](h map[string]any, key string, f func(IN) result.HMapResult[OUT]) result.HMapResult[OUT] {

	v := Get[IN](h, key)
	if !v.Ok() {
		return result.NewMapErrorResult[OUT](v.Err())
	}

	return result.FMap[IN, OUT](v, f)
}

func Get[T any, KEY ~string](h map[KEY]any, key KEY) result.HMapResult[T] {
	v, ok := h[key]
	if !ok {
		return result.NewMapErrorResult[T](NewNotFoundMapError[KEY](key))
	}

	typedV, ok := v.(T)
	if !ok {
		return result.NewMapErrorResult[T](NewInvalidTypeMapError(reflect.TypeOf(new(T)).String(), reflect.TypeOf(v).String()))
	}

	return result.NewMapValidResult(typedV)
}

// to set it, you can do whatever you want
