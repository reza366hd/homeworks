package fourteen

import "fmt"

var t2, t1 int

func fibo(i int) {
	t1 = 0
	t2 = 1

	for o := 0; o < i; o++ {
		t1, t2 = t2, t2+t1
	}
	fmt.Print(t1, "-")
}

func fourteen() {
	fmt.Println("give me fibonachi number")
	var v int
	fmt.Scanln(&v)
	fibo(v - 1)

}
