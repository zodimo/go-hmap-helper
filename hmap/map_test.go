package hmap

import (
	"testing"

	"github.com/zodimo/go-hmap-helper/hmap/result"
)

func TestGet_Valid(t *testing.T) {
	h := map[string]any{"foo": 42}
	res := Get[int](h, "foo")
	if !res.Ok() {
		t.Fatalf("expected Ok, got error: %v", res.Err())
	}
	v, err := res.Value()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 42 {
		t.Errorf("expected 42, got %v", v)
	}
}

func TestGet_NotFound(t *testing.T) {
	h := map[string]any{"foo": 42}
	res := Get[int](h, "bar")
	if res.Ok() {
		t.Fatalf("expected not Ok for missing key")
	}
	if res.Err() == nil {
		t.Errorf("Expected and error, got nil")
	}
	_, ok := res.Err().(HMapKeyNotFoundError)
	if !ok {
		t.Errorf("Expected HMapKeyNotFoundError, got %T", res.Err())
	}

}

func TestGet_InvalidType(t *testing.T) {
	h := map[string]any{"foo": "not-an-int"}
	res := Get[int](h, "foo")
	if res.Ok() {
		t.Fatalf("expected not Ok for invalid type")
	}
	if res.Err() == nil {
		t.Errorf("Expected and error, got nil")
	}
	_, ok := res.Err().(HMapInvalidTypeError)
	if !ok {
		t.Errorf("Expected HMapInvalidTypeError, got %T", res.Err())
	}

}

func TestFMap_Valid(t *testing.T) {
	h := map[string]any{"foo": 2}
	f := func(x int) result.HMapResult[int] { return result.NewMapValidResult(x * 10) }
	res := FMap[int, int](h, "foo", f)
	if !res.Ok() {
		t.Fatalf("expected Ok, got error: %v", res.Err())
	}
	v, _ := res.Value()
	if v != 20 {
		t.Errorf("expected 20, got %v", v)
	}
}

func TestFMap_NotFound(t *testing.T) {
	h := map[string]any{"foo": 2}
	f := func(x int) result.HMapResult[int] { return result.NewMapValidResult(x * 10) }
	res := FMap[int, int](h, "bar", f)
	if res.Ok() {
		t.Fatalf("expected not Ok for missing key")
	}
	if res.Err() == nil {
		t.Errorf("Expected and error, got nil")
	}
	_, ok := res.Err().(HMapKeyNotFoundError)
	if !ok {
		t.Errorf("Expected HMapKeyNotFoundError, got %T", res.Err())
	}

}

func TestFMap_InvalidType(t *testing.T) {
	h := map[string]any{"foo": "not-an-int"}
	f := func(x int) result.HMapResult[int] { return result.NewMapValidResult(x * 10) }
	res := FMap[int, int](h, "foo", f)
	if res.Ok() {
		t.Fatalf("expected not Ok for invalid type")
	}
	_, ok := res.Err().(HMapInvalidTypeError)
	if !ok {
		t.Errorf("Expected HMapInvalidTypeError, got %T", res.Err())
	}
}
