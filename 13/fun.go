package main

import "fmt"

var t2, t1 int

func fun(i int) {
	t1 = 0
	t2 = 1
	for o := 0; o < i; o++ {
		t1, t2 = t2, t2+t1
		fmt.Print(t1, "-")
	}

}

func main() {
	i := 9
	fun(i)

}
