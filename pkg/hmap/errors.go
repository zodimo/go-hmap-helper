package hmap

import "errors"

func NewNotFoundMapError(key string) error {
	return errors.New("Key does not exist in the map: " + key)
}

func NewInvalidTypeMapError(expectedType string, actualType string) error {
	return errors.New("invalid type, expected " + expectedType + " but got " + actualType)
}
