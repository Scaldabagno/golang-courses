# Section 2-4

## Pointers

* Function with pointers to swap values
```go
func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}
```

* Using this method
```go
m1, m2 := 2, 3
fmt.Println(m1, m2) // 2, 3
swap(&m1, &m2) // using their addresses, like C
fmt.Println(m1, m2) // 3, 2
```
