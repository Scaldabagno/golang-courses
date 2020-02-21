# Section 4-4

## Creating our first server

* Import for http functions
```go
import "net/http"
```

* Creates a new server
```go
  mux := http.NewServeMux()
```

* Defines the endpoint and the function to call
```go
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	fmt.Println(r.Method) // print method
	w.Write([]byte("Hello World"))
  })
```

* Starts the server 
```go
	http.ListenAndServe("localhost:3000", mux)
```
