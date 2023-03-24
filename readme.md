# GenerateX

GenerateX is a tool for go code generation.

## Installation

```bash
go get github.com/unsafe-risk/generatex
```

## Usage

### generate

```bash
go generate ./...
```

### Tuple

`//go:generate go run github.com/unsafe-risk/generatex/cmd/tuple <package-name> <tuple-name> <types>...`

```go
package tuple

//go:generate go run github.com/unsafe-risk/generatex/cmd/tuple tuple Pair string float64

//go:generate go run github.com/unsafe-risk/generatex/cmd/tuple tuple Tripple string float64 int64
```

```go
// pair.go
package tuple

type Pair struct {
	V1 string
	V2 float64
}

func NewPair(v1 string, v2 float64) Pair {
	return Pair{
		V1: v1,
		V2: v2,
	}
}

// tripple.go
package tuple

type Tripple struct {
	V1 string
	V2 float64
	V3 int64
}

func NewTripple(v1 string, v2 float64, v3 int64) Tripple {
	return Tripple{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}
```

### Stream

`//go:generate go run github.com/unsafe-risk/generatex/cmd/stream <package-name> <stream-name> <types>...`

```go
package stream

//go:generate go run github.com/unsafe-risk/generatex/cmd/stream stram Stream string int64 int64 float64
```

```go
// stream.go
package stream

type Stream struct {
	F1 func(string) (int64, error)
	F2 func(int64) (int64, error)
	F3 func(int64) (float64, error)
}

func NewStream(f1 func(string) (int64, error), f2 func(int64) (int64, error), f3 func(int64) (float64, error)) *Stream {
	return &Stream{
		F1: f1,
		F2: f2,
		F3: f3,
	}
}

func (s *Stream) Run(init string) (result float64, err error) {
	rs1 := init

	rs2, err := s.F1(rs1)
	if err != nil {
		return result, err
	}

	rs3, err := s.F2(rs2)
	if err != nil {
		return result, err
	}

	rs4, err := s.F3(rs3)
	if err != nil {
		return result, err
	}

	return rs4, nil
}
```

### Union

`//go:generate go run ./cmd/union <package-name> <type-name> <types>...`

```go
package union

//go:generate go run github.com/unsafe-risk/generatex/cmd/union union Union string int64 float64
```

```go
// union.go
package union

type Union struct {
	value any
}

func (m Union) IsInt64() bool {
	_, ok := m.value.(int64)
	return ok
}

func (m Union) AsInt64() (int64, bool) {
	v, ok := m.value.(int64)
	return v, ok
}

func (m Union) MustInt64() int64 {
	v, ok := m.value.(int64)
	if !ok {
		panic("int64 is not the type of the value")
	}
	return v
}

func UnionOfInt64(v int64) Union {
	return Union{
		value: v,
	}
}

func (m Union) IsFloat64() bool {
	_, ok := m.value.(float64)
	return ok
}

func (m Union) AsFloat64() (float64, bool) {
	v, ok := m.value.(float64)
	return v, ok
}

func (m Union) MustFloat64() float64 {
	v, ok := m.value.(float64)
	if !ok {
		panic("float64 is not the type of the value")
	}
	return v
}

func UnionOfFloat64(v float64) Union {
	return Union{
		value: v,
	}
}

func (m Union) IsString() bool {
	_, ok := m.value.(string)
	return ok
}

func (m Union) AsString() (string, bool) {
	v, ok := m.value.(string)
	return v, ok
}

func (m Union) MustString() string {
	v, ok := m.value.(string)
	if !ok {
		panic("string is not the type of the value")
	}
	return v
}

func UnionOfString(v string) Union {
	return Union{
		value: v,
	}
}

```