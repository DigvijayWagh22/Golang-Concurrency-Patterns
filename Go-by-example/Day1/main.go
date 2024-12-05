package main

import (
	"fmt"
	"math"
)

// global const
const s = " constant"

func main() {
	// Basic hello world
	fmt.Println("Hello world")

	// concat the sentence
	fmt.Println("go" + "lang")

	// arithmetic operators
	fmt.Println("2+2 = ", 2+2)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// logical operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// variable declaration
	var a = "initial" // normal
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c) // multivariable declaration

	var d = true // compiler will infer the type
	fmt.Println(d)

	var e int // zero value assigned
	fmt.Println(e)

	f := "digvijay" // short hand notation
	fmt.Println(f)

	// constant

	fmt.Println(s) // Global const

	const n = 500000000
	const n1 = 3e20 / n
	fmt.Println(n1)

	fmt.Println(int64(n1))
	fmt.Println(math.Sin(n))

	// for loop

	i := 1
	for i <= 3 { // basic for loop external value declaration
		fmt.Println(i)
		i++
	}
	fmt.Println()

	for j := 0; j < 3; j++ { // basic for loop with internal declaration
		fmt.Println(j)
	}
	fmt.Println()

	for i := range 3 { // iterates the range number times
		fmt.Println("range ", i)
	}

	for { // break condition
		fmt.Println("will print once and then break")
		break
	}

	for n := range 6 { // continue condition
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// If/Else
	if 7%2 == 0 { // basic if /else
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // basic if
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 { // logical condition based condition
		fmt.Println("Either 7 or 8 are even")
	}

	if num := 9; num < 0 { // If - Else if - else ladder
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num has 1 digit")
	} else {
		fmt.Println("num has multiple digtis")
	}
	// switch statement

}
