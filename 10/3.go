package main

import "fmt"

func main() {
	var count [10000000000]string
	var o string
	for n := 0; n != -1; n++ {
		fmt.Println("Do you have any student")
		fmt.Scanln(&o)
		if o == "yes" {
			fmt.Println("please enter student", n+1, "name")
			fmt.Scanln(&count[n])

		} else {
			break
		}
	}
	fmt.Println(count)
}
