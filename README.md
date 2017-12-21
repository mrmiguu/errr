# Example function (using errr)
```go
func atoi(s string, err ...*error) int {
    n, e := strconv.Atoi(s)

	if errr.Orr(err, e) {
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
    
    if errr.Orr(err) {
        return
    }

    println(n)
}
```