# Section 2-8

## Control/Conditional Structures in Go: Pt. 1

* if
```go
f := true
flag := &f

if flag == nil {
	fmt.Println("Value is nil")
} else if *flag {
	fmt.Println("True")
} else {
	fmt.Println("False")
}

```

* for
```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
} 


// infinite loop
for{

}


// looping inside an array
arr := []string{"dncs", "cnjd", "cndn"}

for i := range arr {
	fmt.Println("Index:", i, "\tValue:", arr[i])
}

// looping better inside an array
for i, value := range arr {
	fmt.Println("Index:", i, "\tValue:", value)
	// fmt.Println(i)
	// fmt.Println(value)
	// fmt.Printf("Index: %v \t Value: %s\t", i, value)
}

// looping inside a map
mymap := make(map[string]interface{})
mymap["name"] = "dnasn"
mymap["age"] = 20

for k, v := range mymap {
	// fmt.Printf("Key: %s \t Value: %v\t", k, v)
	fmt.Println("Key:", k, "\tValue:", v)
}

```

* while does not exist