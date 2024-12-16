package main

import (
	"fmt"
	"unicode/utf8"
)

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name, age: 42}
	return &p
}

func examineRune(r rune) {

	if r == 'H' {
		fmt.Println("found tee")
	} else if r == 'o' {
		fmt.Println("found so sua")
	}
}

type rect struct {
	height  int
	breadth int
}

func (r *rect) Area() int {
	return r.breadth * r.height
}

func (r *rect) Perim() int {
	return 2 * (r.height + r.breadth)
}

func main() {

	// strings and runes
	str := "Hello world"

	fmt.Println("Length of str: ", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
		fmt.Printf("%v ", str[i])
		fmt.Println(str[i])
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(str))

	for idx, value := range str {
		fmt.Printf("idx is %d  value is %#U \n", idx, value)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

	// structs
	fmt.Println(Person{"Bob", 12})
	fmt.Println(Person{name: "Tom", age: 12})
	fmt.Println(Person{})
	fmt.Println(&Person{"Jerry", 21})
	fmt.Println(newPerson("jay"))

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Rex",
		isGood: true,
	}

	fmt.Println(dog)

	s := Person{"Merry", 61}

	sp := &s
	fmt.Println(s.age)
	fmt.Println(s.name)
	fmt.Println(sp.age)
	fmt.Println(sp.name)
	sp.name = "perie"
	sp.age = 65

	fmt.Println(s.age)
	fmt.Println(s.name)
	fmt.Println(sp.age)
	fmt.Println(sp.name)

	// methods

	r := rect{height: 5, breadth: 10}
	fmt.Println(r.Area())
	fmt.Println(r.Perim())

	rp := &r
	fmt.Println(rp.Area())
	fmt.Println(rp.Perim())
}
