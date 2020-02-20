# Section 2-9

## Control/Conditional Structures in Go: Pt. 2

* how to make a while loop
```go
for{
    if condition{
        break
    }
}
```

* continue
```go
flag := true

for i := 0; i < 10; i++ {
	if i%2 == 0 {
		continue // 1 3 5 7 9
	}
	fmt.Println(i)
}
```

* break
```go
flag := true

for i := 0; i < 10; i++ {
	if i%2 == 0 {
		flag = false
		break
	}
	fmt.Println(i) // nothing
}
fmt.Println(flag) // false
```

* switch
```go
// first way
day := "Fri"
switch day {
case "Fri":
	fmt.Println("TGIF")
case "Mon", "Tue", "Wed":
	fmt.Println("dwhua")
default:
	fmt.Println("default")
}

// second way
switch {
case day == "Fri":
	fmt.Println("TGIF")
	break
}

// catching errors
err := errors.New("cbdsuvbsd") // new error
switch err {} // catching errors
```