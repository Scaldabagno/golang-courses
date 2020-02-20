# Section 2-1

## Declaring and initializing variables

### Atomic datatypes

* string
* int
* int32
* int64
* uint
* uint32
* uint64

### Unsafe

* Pointers

### Abstract datatypes

* map[]<datatype>
* struct{}
* interface{}

### Ways to declare variables and else

* Way to declare a variable specifying datatype
```go
var m1 int
m1 = 2
```

* Way to declare more variables
```go
var (
	m1 = 2
	m2 = 3
	)
```

* Summing different datatypes without casting returns an error
```go
var m1 int32
var m2 int64
fmt.Println(m1 + m2) //error
fmt.Println(int64(m1) + m2) //ok
```

* All variables in go, by default are 0

* Another way to declare variables
```go
m1 := 2
m2 := 3
```