package main

import (
	"fmt"
	"time"
)

func main() {
	// switch

	i := 7 // finding a number
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("More that three")
	}

	switch time.Now().Weekday() { // like if else condition for weekday weekend
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its weekday")

	}

	t := time.Now()
	switch {
	case t.Hour() > 12:
		fmt.Println("Its after noon")
	default:
		fmt.Println(" It's before noon")
	}

	switch {
	case time.Now().Hour() > 12:
		fmt.Println("Its after noon")
	default:
		fmt.Println(" It's before noon")
	}

	WhatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know the type  %T\n", t)

		}
	}

	WhatAmI(true)
	WhatAmI(8)
	WhatAmI("hello")

	// Arrays
	var a [5]int // normal declaration
	fmt.Println("emp", a)

	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	fmt.Println("get length ", len(a))

	var twoD1 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD1[i][j] = i + j
		}
	}
	fmt.Println("twoD1 : ", twoD1)

	twoD2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("twoD2 : ", twoD2)

	b := [5]int{1, 2, 3, 4, 5} // short hand notation
	fmt.Println("dcl : ", b)

	b = [...]int{6, 7, 4, 5, 6} // let compiler count the size
	fmt.Println("dcl : ", b)

	b = [...]int{1, 2: 2, 3, 4}
	fmt.Println("zero  values :", b)

}
