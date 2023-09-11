package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := map[string]int{}

	var result map[string]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = Sentence(line, data)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func Sentence(Sentence string, data map[string]int) map[string]int {
	counter := 0
	for i := range Sentence {
		if string(Sentence[i]) == " " {
			word := Sentence[counter:i]
			data[word]++
			counter = i + 1
		}
	}
	lastWord := Sentence[counter:]
	data[lastWord]++
	return data
}
