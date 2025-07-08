package hmap

import (
	"reflect"
)

type HMapKey[T any] = string

type HMap[T any] interface {
	GetWithKey(key HMapKey[T]) HMapResult[T]
}

var _ HMap[any] = &hmap[any]{}

type hmap[T any] struct {
	m map[string]any
}

func (h *hmap[T]) GetWithKey(key HMapKey[T]) HMapResult[T] {
	v, ok := h.m[key]
	if !ok {
		return NewMapErrorResult[T](NewNotFoundMapError[T](key))
	}

	typedV, ok := v.(T)
	if !ok {
		return NewMapErrorResult[T](NewInvalidTypeMapError[T](reflect.TypeOf(new(T)).String(), reflect.TypeOf(v).String()))
	}

	return NewMapValidResult(typedV)
}
