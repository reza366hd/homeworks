package nineteen

import (
	"fmt"
)

func aval(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func nineteen() {

	var shif [2]int
	var i int
	var p int
	for i := 0; i < 100; i++ {
		if aval(i) {
			if i < 10 {
				fmt.Println(i)
			}
			if 10 < i && i < 100 {

				if aval(p) {
					fmt.Println(p)

				}

			}

		}

	}

	fmt.Println(i)
	o := 48
	for j := 0; true; j++ {

		shif[j] = o % 10
		if o >= 10 {
			o = o / 10
		} else {
			break
		}
	}
	p = shif[0]*10 + shif[1]
	fmt.Println(p)

}
