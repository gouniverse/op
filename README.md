# op

## Installation

```
go get github.com/gouniverse/op
```

## Usage

- IfElse

```go
display := op.IfElse(str == "pass", "passed", "failed")
```

- IfNil

```go
display := op.IfNil(displayMessage, "failed")
```

- IfNilOrEmpty

```go
display := op.IfNilOrEmpty(displayMessage, "failed")
```

- IfNilOrWhitespace

```go
display := op.IfNilOrWhitespace(displayMessage, "failed")
```

- NilCoalescing (same as IfNil)

```go
display := op.NilCoalescing(displayMessage, "failed")
```

- Ternary (same as IfElse)

```go
display := op.Ternary(str == "pass", "passed", "failed")
```
