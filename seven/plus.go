package seven

import "fmt"

func seven() {

	first := 1
	end := 6
	H := 0

	for i := first; i <= end; i++ {
		H = i + H
		fmt.Println(H)
	}
	fmt.Println("out_side", H)
}
