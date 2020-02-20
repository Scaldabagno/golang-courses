package main

import "fmt"

func main() {

	flag := true

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// continue // 1 3 5 7 9
			flag = false
			break // false
		}
		fmt.Println(i)
	}
	fmt.Println(flag)

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Tue", "Wed":
		fmt.Println("dwhua")
	default:
		fmt.Println("default")
	}

	switch {
	case day == "Fri":
		fmt.Println("TGIF")
		break
	}

	// err := errors.New("cbdsuvbsd") // new error
	// switch err {} // catching errors

}
