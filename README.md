# anycase

[![Godoc Reference](https://godoc.org/github.com/loopcontext/anycase?status.svg)](http://godoc.org/github.com/loopcontext/anycase)
[![Build Status](https://travis-ci.org/loopcontext/anycase.svg)](https://travis-ci.org/loopcontext/anycase)
[![Coverage](http://gocover.io/_badge/github.com/loopcontext/anycase?0)](http://gocover.io/github.com/loopcontext/anycase)
[![Go Report Card](https://goreportcard.com/badge/github.com/loopcontext/anycase)](https://goreportcard.com/report/github.com/loopcontext/anycase)

anycase is a go package for converting string case to various cases (e.g.
[snake case](https://en.wikipedia.org/wiki/Snake_case) or
[camel case](https://en.wikipedia.org/wiki/CamelCase)) or
[kebab case](https://en.wiktionary.org/wiki/kebab_case)) to see the full conversion table below.

## Install

```bash
go get -u github.com/loopcontext/anycase
```

## Example

```go
s := "AnyKind of_string"
```

| Function                                  | Result               |
| ----------------------------------------- | -------------------- |
| `ToSnake(s)`                              | `any_kind_of_string` |
| `ToSnakeAndIgnore(s, '.')`                | `any_kind.of_string` |
| `ToSnakeUppercase(s)`                     | `ANY_KIND_OF_STRING` |
| `ToKebab(s)`                              | `any-kind-of-string` |
| `ToKebabUppercase(s)`                     | `ANY-KIND-OF-STRING` |
| `ToCamel(s)`                              | `AnyKindOfString`    |
| `ToLowerCamel(s)`                         | `anyKindOfString`    |
| `ToDelimited(s, '.')`                     | `any.kind.of.string` |
| `ToDelimitedUppercase(s, '.', ' ', true)` | `ANY.KIND OF.STRING` |
| `ToDelimitedUppercase(s, '.', '', true)`  | `ANY.KIND.OF.STRING` |
