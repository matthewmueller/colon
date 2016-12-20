# Colon

Little colon templating engine written in Go.

This is a port of it's [node.js library](https://github.com/matthewmueller/colon-template).

## Example

```go
render := Compile("hi :name")
s := render(map[string]interface{}{"name": "matt"})
// hi matt
```

## Installation

```
go get github.com/matthewmueller/colon
```
