package ten

import (
	"fmt"
	"strconv"
)

func ten() {
	var first string
	var second string
	var plus string
	fmt.Scanln(&first)
	fmt.Scanln(&second)
	fmt.Scanln(&plus)

	intfirst, _ := strconv.ParseInt(first, 10, 0)
	intsecond, _ := strconv.ParseInt(second, 10, 0)

	if plus == "/" {
		fmt.Printf("= %d \n", intfirst/intsecond)
	} else if plus == "*" {
		fmt.Printf("= %d \n", intfirst*intsecond)
	} else if plus == "-" {
		fmt.Printf("= %d \n", intfirst-intsecond)
	} else if plus == "+" {
		fmt.Printf("= %d \n", intfirst+intsecond)
	} else {
		fmt.Println("eror")
	}
}
