package main

import (
	"fmt"
	"time"
)

func main() {

	// Normal variable declaration
	var fname string
	fname = "Aditya"
	fmt.Println(fname)

	// Short hand variable declaration
	lname := "Bisht"
	fmt.Println(lname)

	// Multiple variable declaration
	var age int = 22
	fmt.Println(age)

	// Constant variable declaration
	const isMarried bool = false
	fmt.Println(isMarried)

	// Float variable declaration
	const height float64 = 5.11
	fmt.Println(height)
	var area = 3.14 * height * height
	fmt.Println(area)

	// for loop variable declaration
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	// range
	for i := 0; i < 3; i++ {
		fmt.Println(i + 1)
	}

	//switch case variable declaration
	i = 3
	switch i {
	case 1:
		fmt.Println("i is 1")
	default:
		fmt.Println("i is not 1")
	}

	// multiple constant switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//Type Switch
	whoAmI := "Aditya"
	switch interface{}(whoAmI).(type) {
	case string:
		fmt.Println("I am a string")
	default:
		fmt.Println("I am not a string")
	}

}
