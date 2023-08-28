package eight

import "fmt"

func eight() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

}
