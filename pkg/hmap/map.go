package hmap

import (
	"reflect"
)

type HMapKey[T any] = string

type HMap[T any] interface {
	GetWithKey(key HMapKey[T]) HMapResult[T]
	Unwrap() map[string]any
}

var _ HMap[any] = &hmap[any]{}

type hmap[T any] struct {
	m map[string]any
}

func (h *hmap[T]) GetWithKey(key HMapKey[T]) HMapResult[T] {
	v, ok := h.m[string(key)]
	if !ok {
		return NewMapErrorResult[T](NewNotFoundMapError(key))
	}

	typedV, ok := v.(T)
	if !ok {
		return NewMapErrorResult[T](NewInvalidTypeMapError(reflect.TypeOf(new(T)).String(), reflect.TypeOf(v).String()))
	}

	return NewMapValidResult(typedV)
}

func (h *hmap[T]) Unwrap() map[string]any {
	return h.m
}

func FMap[IN, OUT any](h HMap[any], key string, f func(IN) HMapResult[OUT]) HMapResult[OUT] {

	v, ok := h.Unwrap()[key]
	if !ok {
		return NewMapErrorResult[OUT](NewNotFoundMapError(key))
	}
	typedV, ok := v.(IN)
	if !ok {
		return NewMapErrorResult[OUT](NewInvalidTypeMapError(reflect.TypeOf(new(IN)).String(), reflect.TypeOf(v).String()))
	}

	return f(typedV)
}
