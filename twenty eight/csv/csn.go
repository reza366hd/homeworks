package main

import (
	"encoding/csv"
	"os"
)

func main() {

	data := [][]string{
		{"first_name", "last_name"},
		{"mamad", "dvd"},
		{"reza", "mi"},
		{"bob", "sfanji"},
	}

	f, _ := os.Create("users.csv")
	defer f.Close()

	w := csv.NewWriter(f)
	_ = w.WriteAll(data)
}
