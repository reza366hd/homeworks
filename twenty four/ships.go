package twentyfour

import (
	"fmt"
)

func twentyfour() {
	a := "babao"

	var o ships
	o.Name = "maw"

	o.flock(a)

	fmt.Println(o)
}

func (h ships) flock(a string) {
	h.Next = new(ships)
	h.Next.Name = a
}

type ships struct {
	Name string
	Next *ships
}
