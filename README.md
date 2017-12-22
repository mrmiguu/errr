Example function (using errr)
```go
func atoi(s string, e ...*error) int {
    n, err := strconv.Atoi(s)

    if errr.Is(e, err) {
        return 0
    }
    return n
}
```
Panic on error
```go
n := atoi("123")
println(n)
```
Error handling
```go
var err error

n := atoi("123", &err)
if err != nil {
    return err
}

println(n)
```
Error handling (using errr)
```go
func test(e ...*error) {
    n := atoi("123", e...)

    if errr.Is(e) {
        return
    }

    println(n)
}
```