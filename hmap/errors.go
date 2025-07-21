package hmap

import "fmt"

var _ error = (*HMapKeyNotFoundError)(nil)
var _ error = (*HMapInvalidTypeError)(nil)

type HMapKeyNotFoundError struct {
	Key string
}

func (e HMapKeyNotFoundError) Error() string {
	return fmt.Sprintf("key[%s] does not exist in the map", e.Key)
}

type HMapInvalidTypeError struct {
	ExpectedType string
	ActualType   string
}

func (e HMapInvalidTypeError) Error() string {
	return fmt.Sprintf("invalid type, expected %s but got %s", e.ExpectedType, e.ActualType)
}

func NewNotFoundMapError[KEY ~string](key KEY) HMapKeyNotFoundError {
	return HMapKeyNotFoundError{Key: string(key)}
}

func NewInvalidTypeMapError(expectedType string, actualType string) HMapInvalidTypeError {
	return HMapInvalidTypeError{ExpectedType: expectedType, ActualType: actualType}
}
