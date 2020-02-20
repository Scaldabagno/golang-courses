# Section 2-7

## Interfaces: Pt. 2

* Function that takes anything in input
```go
func Anything(anything interface{}) {
	fmt.Println(anything)
}

// Usage
Anything(2.44)       // 2.44
Anything("angad")    // angad
Anything(struct{}{}) // {}
```

* Making functions with map (~lambda)
```go
mymap := make(map[string]interface{})

mymap["name"] = "nIUHSHBCI"
mymap["age"] = 10
fmt.Println(mymap)
// map[age:10 name:nIUHSHBCI]
```
