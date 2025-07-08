package hmap

type HMapResult[T any] struct {
	Value T
	Ok    bool
	Err   error
}

func NewMapErrorResult[T any](err error) HMapResult[T] {
	return HMapResult[T]{
		Value: *new(T),
		Ok:    false,
		Err:   err,
	}
}

func NewMapValidResult[T any](value T) HMapResult[T] {
	return HMapResult[T]{
		Value: value,
		Ok:    true,
		Err:   nil,
	}
}
