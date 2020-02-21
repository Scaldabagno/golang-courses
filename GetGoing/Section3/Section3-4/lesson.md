# Section 3-4

## Goroutines: Pt. 1

* Goroutines are a way to call a function in background
```go
heavy() // normal call
go heavy() // goroutine -> way to call a function, background
go superHeavy() // as a goroutine, it does not block main execution
fmt.Println("End")
// time.Sleep(time.Second * 5) // time for running goroutine (bad way)
select {} // right way
```

* Goroutines are executed concurrently
