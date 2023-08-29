package ananymos

import "fmt"

func ananymos() {

	var persen []struct {
		height uint8
		Name   string
		Age    uint8
	} = []struct {
		height uint8
		Name   string
		Age    uint8
	}{
		{Name: "pat", Age: 19, height: 150},
		{Name: "reza", Age: 16, height: 170},
		{Name: "mamad", Age: 22, height: 182},
	}

	fmt.Println(persen)
}
