# Section 3-5

## Goroutines: Pt. 2

* Sync package
```go
import "sync"
```

```go
// mutexes
mut := &sync.Mutex{}
mut.Lock()
mut.Unlock()
```

* Concurrent programming
```go
// wait groups
wg := &sync.WaitGroup{}
// add, done, wait
wg.Add(2) // adds 2 funcs, after 2 done calls, it will go on
// lambda/anonymous functions as goroutines
go func() {
  fmt.Println("Hello")
  wg.Done() // close
}()
go func() {
  fmt.Println("World")
  wg.Done() // close
}()

wg.Wait() // unlock

fmt.Println("End")
// select {} // deadlock

```
