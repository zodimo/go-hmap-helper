package result

type HMapResult[T any] struct {
	value T
	ok    bool
	err   error
}

func NewMapErrorResult[T any](err error) HMapResult[T] {
	return HMapResult[T]{
		value: *new(T),
		ok:    false,
		err:   err,
	}
}

func NewMapValidResult[T any](value T) HMapResult[T] {
	return HMapResult[T]{
		value: value,
		ok:    true,
		err:   nil,
	}
}

func (r HMapResult[T]) Value() (T, error) {
	if r.err != nil {
		return *new(T), r.err
	}

	return r.value, nil
}

func (r HMapResult[T]) ValueOr(defaultValue T) T {
	if !r.ok {
		return defaultValue
	}

	return r.value
}

func (r HMapResult[T]) Ok() bool {
	return r.ok
}

func (r HMapResult[T]) Err() error {
	return r.err
}

func FMap[IN, OUT any](r HMapResult[IN], f func(IN) HMapResult[OUT]) HMapResult[OUT] {
	if !r.ok {
		return NewMapErrorResult[OUT](r.err)
	}

	return f(r.value)
}
