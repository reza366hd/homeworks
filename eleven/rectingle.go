package eleven

import "fmt"

func eleven() {
	var width, height int

	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&width)

	fmt.Print("Enter the height of the rectangle: ")
	fmt.Scan(&height)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
