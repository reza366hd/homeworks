package eight

import "fmt"

func eight() {
	var e int
	for i := 0; i < 20; i++ {
		if i%3 == 0 {
			e += i
		}
		if i%5 == 0 {
			e += i
		}
	}
	fmt.Println(e)
}
