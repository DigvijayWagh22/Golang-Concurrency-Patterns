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

}
