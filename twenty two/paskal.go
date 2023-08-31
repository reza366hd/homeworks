package twentytwo

import "fmt"

func twentytwo() {
	n := 7

	arr := make([]int, n)

	for i := range arr {
		arr[i] = 0
	}

	end := len(arr)
	arr[end-1] = 1

	for i := 0; i < n*2; i++ {
		for j := range arr {
			fmt.Print(arr[j])
		}

		paskal(&arr)

		fmt.Println()
	}
}

func paskal(arrPointer *[]int) {
	found := false
	var k int = 1
	for i := range *arrPointer {
		if (*arrPointer)[i] == k && !found {

			(*arrPointer)[i-1] = k
			found = true
		}
	}
}
