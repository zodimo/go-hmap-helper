# go-typed-maps/hmap
A collection of generic, type-safe helpers for working with Go maps using string keys and `any` values. This package provides robust error handling and functional-style utilities for extracting and transforming values from maps, with a focus on type safety and composability.

## Features
- Type-safe extraction from `map[string]any` with error reporting
- Functional mapping (`FMap`) over map values
- Composable result types for error handling
- Custom error types for missing keys and type mismatches

## Installation

Add to your Go project:

```sh
go get github.com/zodimo/go-hmap-helper
```

## Usage

### Basic Extraction

```go
package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
)

func main() {
	h := map[string]any{"foo": 42}
	res := hmap.Get[int](h, "foo")
	if res.Ok() {
		v, _ := res.Value()
		fmt.Println(v) // 42
	} else {
		fmt.Println("Error:", res.Err())
	}
}

```

### Handling Missing Keys

```go
package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
)

func main() {
	h := map[string]any{"foo": 42}
	res := hmap.Get[int](h, "bar")
	if !res.Ok() {
		// res.Err() will be HMapKeyNotFoundError
		fmt.Println(res.Err())
	}
}

```

### Handling Type Errors

```go
package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
)

func main() {
	h := map[string]any{"foo": "not-an-int"}
	res := hmap.Get[int](h, "foo")
	if !res.Ok() {
		// res.Err() will be HMapInvalidTypeError
		fmt.Println(res.Err())
	}
}

```

### Functional Mapping with FMap

```go
package main

import (
	"fmt"

	"github.com/zodimo/go-hmap-helper/hmap"
	"github.com/zodimo/go-hmap-helper/hmap/result"
)

func main() {
	h := map[string]any{"foo": 42}
	f := func(x int) result.HMapResult[int] { return result.NewMapValidResult(x * 10) }
	res := hmap.FMap[int, int](h, "foo", f)
	if res.Ok() {
		v, _ := res.Value()
		fmt.Println(v) // 420 if h["foo"] == 42
	}
}
```

## API Reference

### `Get[T any](h map[string]any, key string) HMapResult[T]`
Safely extract a value of type `T` from the map. Returns a result type that can be checked for success or error.

### `FMap[IN, OUT any](h map[string]any, key string, f func(IN) HMapResult[OUT]) HMapResult[OUT]`
Extracts a value and applies a function to it if successful, returning a new result.

### `HMapResult[T any]`
A result type representing either a valid value or an error. Methods:
- `Ok() bool` — true if the value is valid
- `Err() error` — returns the error if present
- `Value() (T, error)` — returns the value or error
- `ValueOr(defaultValue T) T` — returns the value or a default

### Error Types
- `HMapKeyNotFoundError` — for missing keys
- `HMapInvalidTypeError` — for type mismatches

## Testing

Run tests with:

```sh
go test ./...
```

## License

MIT 