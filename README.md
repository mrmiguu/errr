# Example function (using errr)
```go
func atoi(s string, err ...*error) int {
	e := errr.New(err...)

	n, or := strconv.Atoi(s)
	if e.Err(or) {
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

# Error handling (using errr)
```go
func test(err ...*error) {
    n := atoi("123", err...)
    
    if errr.Is(err...) {
        return
    }

    println(n)
}
```