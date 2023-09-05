package main

import (
	"fmt"
	ds "qurue/queue"
)

func main() {
	s := ds.Queue{}

	s.Push(10)
	s.Push(20)
	s.Print()
	fmt.Println(s.Pop())
	s.Print()
	fmt.Println(s.Read(0))

}
