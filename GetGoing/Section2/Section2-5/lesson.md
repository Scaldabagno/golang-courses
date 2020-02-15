# Section 2-5

## Structs

* Declaring a struct
```go
type Car struct {
	Name     string
	Age      int
	ModelnNo int
}
```

* Defining an object through a struct
```go
c := Car{
	Name:     "ah",
	Age:      1,
	ModelnNo: 2,
}
```

* Accessing variables of an object
```go
fmt.Println(c.getName()) // ah
```
