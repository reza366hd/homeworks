package main

import (
	"fmt"
)

func aval(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func finder(had int) int {
	primenumss := []int{}
	for i := 2; i < had; i++ {
		if aval(i) {
			primenumss = append(primenumss, i)
		}
	}

	var jam int
	var tedad int

	for i := 0; i < len(primenumss); i++ {
		sum := 0
		count := 1
		for j := i; j < len(primenumss); j++ {
			sum += primenumss[j]
			count++
			if sum >= had {
				break
			}
			if aval(sum) && count > tedad {
				jam = sum
				tedad = count
			}
		}
	}

	return jam
}

func main() {
	had := 1000 //000
	result := finder(had)
	fmt.Println(result)
}
