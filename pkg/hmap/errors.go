package hmap

import "errors"

func NewNotFoundMapError[T any](key string) error {
	return errors.New("Key does not exist in the map: " + key)
}

func NewInvalidTypeMapError[T any](expectedType string, actualType string) error {
	return errors.New("invalid type, expected " + expectedType + " but got " + actualType)
}
