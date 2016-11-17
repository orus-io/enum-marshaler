[![Go Report Card](https://goreportcard.com/badge/github.com/orus-io/enum-marshaler)](https://goreportcard.com/report/github.com/orus-io/enum-marshaler)


# Enum Marshaler

Enum-marshaler is a tool to automate the creation of methods that satisfy the
encoding.TextMarshaler, encoding.TextUnmarshaler, and fmt.Stringer interfaces.

TextMarshaler and TextUnmarshaler are notably used by json or yaml encodings,
which makes it easy and natural to use enums in the code with string representation
in those encodings.

The initial implementation is a copy of
https://godoc.org/golang.org/x/tools/cmd/stringer to which we added the generation
of MarshalText and UnmarshalText, as well as an option to strip prefixes from
constant names in (un)marshaling functions.

## Install

```bash
get get -u github.com/orus-io/enum-marshaler
get install github.com/orus-io/enum-marshaler
```

## Usage

Usage is the same as stringer, with only one extra option, and more code generated.

Given this snippet:

```go
package painkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
```

Running this command:

```bash
enum-marshaler -type Pill
```

in the same directory will create the file `pill_string.go`, in package `painkiller`,
containing a definition of

```go
func (Pill) String() string
func (Pill) MarshalText() []byte, err
func (*Pill) UnmarshalText([]byte) err
```

Typically this process would be run using go generate, like this:

```go
//go:generate enum-marshaler -type=Pill
```
