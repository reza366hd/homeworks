package main

import (
	"fmt"
)

func main() {
	Sentence := "hi reza hi hi hi bro joon "

	var data map[string]int
	data = map[string]int{}
	var repositori string

	counter := 0
	for i := range Sentence {

		if string(Sentence[i]) == " " || string(Sentence[i]) == "" {

			repositori = (string(Sentence[counter:i]))
			data[repositori] += 1
			counter = 0
			counter += i + 1

		}
	}
	fmt.Println(data)
}
