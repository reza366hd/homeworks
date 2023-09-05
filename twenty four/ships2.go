package main

import (
	"fmt"
)

func main() {
	sh := flock{Name: "babao"}

	sh.Next = new(flock)

	sh.Next.Name = "baba"

	sh.Next.Next = new(flock)

	sh.Next.Next.Name = "babaey"

	sh.Next.Next.Next = new(flock)

	sh.Next.Next.Next.Name = "maaaa"

	fmt.Println(sh)

}

type flock struct {
	Name string
	Next *flock
}
