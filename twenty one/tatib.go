package twentyone

import "fmt"

func twentyone() {
	arr := []int{5, 8, 4, 9}
	tatib := tartib(arr)
	fmt.Println(tartib(tatib))
}

func tartib(arr []int) []int {
	var j int
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j = 0; j < n-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}

	return arr
}
