# Example function (using errr)
```go
func atoi(s string, opt ...*error) int {
	e := errr.New(opt...)

	n, err := strconv.Atoi(s)
	if err != nil {
		e.Set(err)
		return 0
	}

	return n
}
```

# Panic on error
```go
n := atoi("123")
println(n)
```

# Error handling
```go
var err error

n := atoi("123", &err)
if err != nil {
    return err
}

println(n)
```