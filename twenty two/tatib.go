package twentytwo

import "fmt"

func twentytwo() {
	arr := []int{4, 8, 1, 5}
	tatib := tartib(arr)
	fmt.Println(tartib(tatib))
}

func tartib(arr []int) []int {

	var j int
	var t int
	var k int
	n := len(arr)
	for j = 1; j < n; j++ {
		t = 1
		k = 0
		for {
			if j-t < 0 || j-k < 0 {
				break
			} else if arr[j-k] < arr[j-t] {
				arr[j-k], arr[j-t] = arr[j-t], arr[j-k]
				k++
				t++
			} else {
				break
			}

		}
	}

	return arr
}
