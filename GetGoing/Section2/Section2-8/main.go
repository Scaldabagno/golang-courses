package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	f := true
	flag := &f

	// no while
	if flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// infinite loop
	// for{
	// }

	arr := []string{"dncs", "cnjd", "cndn"}

	// for i := range arr {
	// 	fmt.Println("Index:", i, "\tValue:", arr[i])
	// }

	for i, value := range arr {
		fmt.Println("Index:", i, "\tValue:", value)
		// fmt.Println(i)
		// fmt.Println(value)
		// fmt.Printf("Index: %v \t Value: %s\t", i, value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "dnasn"
	mymap["age"] = 20

	for k, v := range mymap {
		// fmt.Printf("Key: %s \t Value: %v\t", k, v)
		fmt.Println("Key:", k, "\tValue:", v)
	}
}
