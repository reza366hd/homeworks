package eighteen

import "fmt"

func eighteen() {
	n := 19

	arr := make([]string, n)

	for i := range arr {
		arr[i] = " "
	}

	mid := n / 2
	arr[mid] = "*"

	for i := 0; i < n*2-1; i++ {
		for j := range arr {
			fmt.Print(arr[j])
		}
		if i < mid {
			IncreaseStars(&arr)
		} else {
			DecreaseStars(&arr)
		}
		fmt.Println()
	}
}

func IncreaseStars(arrPointer *[]string) {
	found := false

	for i := range *arrPointer {
		if (*arrPointer)[i] == "*" && !found {
			(*arrPointer)[i-1] = "*"
			found = true
		}

		if (*arrPointer)[i] == " " && found {
			(*arrPointer)[i] = "*"
			found = false
		}
	}
}

func DecreaseStars(arrPointer *[]string) {
	found := false
	for i := range *arrPointer {
		if (*arrPointer)[i] == "*" && !found {
			(*arrPointer)[i] = " "
			found = true
			continue
		}

		if ((*arrPointer)[i] == " " || i == (len(*arrPointer)-1)) && found {
			if i == (len(*arrPointer)-1) && (*arrPointer)[i] == "*" {
				(*arrPointer)[i] = " "
			} else {
				(*arrPointer)[i-1] = " "
			}
			found = false
		}
	}
}
