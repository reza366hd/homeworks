package fourteen

import "fmt"

var i int

func fourteen() {
	for i = 1; true; i++ {
		g := true
		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				g = false
			}
		}
		if g == true {
			fmt.Println(i)
			break
		}
	}
}
