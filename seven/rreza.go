package seven

import "fmt"

func seven() {
	score := 2
	for i := 1; i < 1000; i = i*i + 1 {
		fmt.Println("score : ", score)
		fmt.Println("i : ", i)
		score *= i
		fmt.Println("score * i = ", score)
	}
}
