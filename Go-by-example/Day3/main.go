package main

import (
	"fmt"
	"maps"
	"slices"
)

func plus(a, b int) int {
	return a + b

}

func plusplus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func closurefunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	// slices
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	s = append(s, "e", "f", "g")
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("s1: ", l)

	l = s[2:]
	fmt.Println("s2: ", l)

	l = s[2:]
	fmt.Println("s3: ", l)

	t1 := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}

	if slices.Equal(t1, t2) {
		fmt.Println("t1 == t2")
	}

	twoD := make([][]int, 3)

	for i := 0; i < len(twoD); i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}

	}

	fmt.Println(twoD)

	//maps

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println(len(m))
	fmt.Println(m)

	delete(m, "k2")
	fmt.Println("delete: ", m)

	// clear(m)
	// fmt.Println("clear ", m)

	value, flag := m["k1"]
	fmt.Println("flag : ", flag, "value : ", value)

	n1 := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n1, n2) {
		fmt.Println("n1 == n2")
	}

	// golang functions
	ans1 := plus(1, 2)
	ans2 := plusplus(1, 2, 3)
	fmt.Println(ans1, " ", ans2)

	// multiple return values
	a, b := vals()
	fmt.Println(a, " ", b)
	fmt.Println(vals())

	q, _ := vals()
	fmt.Println(q)

	// Variadic functions
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5)

	// closures
	closefuncvalue := closurefunc()
	newclosefuncvalue := closurefunc()
	fmt.Println(closefuncvalue())
	fmt.Println(closefuncvalue())
	fmt.Println(closefuncvalue())
	fmt.Println(closefuncvalue())
	fmt.Println(closefuncvalue())
	fmt.Println(newclosefuncvalue())
	fmt.Println(newclosefuncvalue())

	// recursion

	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

}
