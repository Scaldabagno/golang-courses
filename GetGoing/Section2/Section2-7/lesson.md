# Section 2-6

## Interfaces: Pt. 2

* Declaring an interface
```go
type Car interface {
	Drive()
	Stop()
}
```

* Defining some structs and creating a relation between them and an interface
```go
type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

// forces to declare all the methods of the interface to a specific structure
func newModel(arg string) Car {
	return &Lambo{arg}
}

// Drive/stop methods defined for structs....
```
