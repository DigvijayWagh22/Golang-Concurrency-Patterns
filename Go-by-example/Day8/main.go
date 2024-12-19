// errors are related to iter package
package main

import (
	"fmt"
	"slices"

	"iter"
)

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type element[T any] struct {
	next *element[T]
	val  T
}
type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	// Generics
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo: ", SliceIndex(s, "zoo"))

	_ = SliceIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list: ", lst.AllElements())

	// Range over iterators
	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
