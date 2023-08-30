package ananymos

import "fmt"

func ananymos() {
	fmt.Println("enter the number")
	var num int
	fmt.Scanln(&num)
	fmt.Println("enter the oprator(enter * to times two // enter ** to the power of two)")
	var oprator string
	fmt.Scanln(&oprator)

	Power := HigherFunction()

	fmt.Println(Power(num, oprator))
}

func HigherFunction() func(i int, b string) int {
	return func(i int, b string) int {
		if b == "**" {
			return i * i
		} else if b == "*" {
			return i + i
		} else {
			return 0
		}
	}
}

func Second(f func() int) int {
	return f()
}
