package nineteen

import "fmt"

func nineteen() {

	name := "reza"
	fmt.Println(name)

	pointer := &name
	fmt.Println(pointer)
	fmt.Println(*pointer)

	pointerpointer := &pointer
	fmt.Println(pointerpointer)
	fmt.Println(&pointerpointer)
	fmt.Println(*&pointerpointer)
	fmt.Println(**pointerpointer)
}
