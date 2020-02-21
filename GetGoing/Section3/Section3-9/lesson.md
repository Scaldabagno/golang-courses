# Section 3-9

## Pong with concurrency 

* Ping
```go
// Pinger pings and waits for pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Millisecond * 500)
		ponger <- 1
	}
}
```

* Pong
```go
// Ponger pongs and waits for ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Millisecond * 500)
		pinger <- 1
	}
}
```

* Both ping and pong are int channels
```go
func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// first ping to make all start
	ping <- 1

	// loop goes forever: ping - pong - ping - pong - etc. //bad
	// for {
	// 	time.Sleep(time.Second)
	// }

	// loop goes forever: ping - pong - ping - pong - etc. //right
	select {}
```
