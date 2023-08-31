package main

import (
	"fmt"
	"mmgi/ds"
)

func main() {
	s := ds.Stack{}

	s.Push(10)
	s.Push(20)
	s.Print()
	fmt.Println(s.Pop())
	s.Print()
	fmt.Println(s.Read(0))

}
