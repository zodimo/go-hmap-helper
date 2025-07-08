package result

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewMapErrorResult(t *testing.T) {
	err := errors.New("test error")
	res := NewMapErrorResult[int](err)
	if res.Ok() {
		t.Errorf("expected Ok to be false")
	}
	if res.Err() != err {
		t.Errorf("expected error to be set")
	}
	v, e := res.Value()
	if e == nil {
		t.Errorf("expected error from Value()")
	}
	if v != 0 {
		t.Errorf("expected zero value, got %v", v)
	}
	if res.ValueOr(42) != 42 {
		t.Errorf("expected ValueOr to return default value")
	}
}

func TestNewMapValidResult(t *testing.T) {
	res := NewMapValidResult(123)
	if !res.Ok() {
		t.Errorf("expected Ok to be true")
	}
	if res.Err() != nil {
		t.Errorf("expected error to be nil")
	}
	v, e := res.Value()
	if e != nil {
		t.Errorf("unexpected error from Value(): %v", e)
	}
	if v != 123 {
		t.Errorf("expected value 123, got %v", v)
	}
	if res.ValueOr(42) != 123 {
		t.Errorf("expected ValueOr to return value, got %v", res.ValueOr(42))
	}
}

func TestFMap(t *testing.T) {
	// FMap with Ok result
	res := NewMapValidResult(5)
	f := func(x int) HMapResult[string] {
		return NewMapValidResult(fmt.Sprintf("val:%d", x))
	}
	mapped := FMap(res, f)
	if !mapped.Ok() {
		t.Errorf("expected Ok to be true")
	}
	v, e := mapped.Value()
	if e != nil {
		t.Errorf("unexpected error: %v", e)
	}
	if v != "val:5" {
		t.Errorf("expected 'val:5', got %v", v)
	}

	// FMap with error result
	err := errors.New("fail")
	errRes := NewMapErrorResult[int](err)
	mappedErr := FMap(errRes, f)
	if mappedErr.Ok() {
		t.Errorf("expected Ok to be false")
	}
	if mappedErr.Err() != err {
		t.Errorf("expected error to be propagated")
	}
}
