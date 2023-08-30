package sixteen

import "fmt"

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

func sixteen() {
	var n []int
	var end int
	fmt.Println("enter number of prime prime number:  ")
	fmt.Scanln(&end)

	for i := 1; true; i++ {
		if aval(i) {
			n = append(n, i)
		}
		if len(n) == end {
			fmt.Print(i)
			break
		}
	}
}
