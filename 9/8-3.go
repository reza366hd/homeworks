package nine

import "fmt"

func nine() {
	var e int
	for i := 0; i < 20; i++ {
		if i%3 == 0 {
			e += i
		} else if i%5 == 0 {
			e += i
		}
		fmt.Println("i =", i)
		fmt.Println()
		fmt.Println("e =", e)
	}
	fmt.Println("outside: ", e)
}
