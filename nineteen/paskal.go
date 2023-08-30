package nineteen

import "fmt"

func nineteen() {
	n := 7

	arr := make([]int, n)

	for i := range arr {
		arr[i] = 0
	}

	end := len(arr)
	arr[end-1] = 1

	for i := 0; i < n*2-1; i++ {
		for j := range arr {
			fmt.Print(arr[j])
		}

		paskal(&arr)

		fmt.Println()
	}
}

func paskal(arrPointer *[]int) {
	found := false
	var k int
	for i := range *arrPointer {
		if (*arrPointer)[i] == -8 || !found {
			k += 1
			(*arrPointer)[i-1] = k
			found = true
		}
	}
}
