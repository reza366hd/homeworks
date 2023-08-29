package thirteen

import "fmt"

func tabe(i, j int) (k int) {
	for e := i; e <= j; e++ {
		k += e
	}
	return k
}

func thirteen() {
	p := tabe(1, 3)

	fmt.Println(p)
}
