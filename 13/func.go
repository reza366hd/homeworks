package main

import "fmt"

func tabe(i, j int) (k int) {
	for e := i; e <= j; e++ {
		k += e
	}
	return k
}

func main() {
	p := tabe(1, 3)

	fmt.Println(p)
}
