package main

import (
	"encoding/json"
	"io/ioutil"
)

type tshirt struct {
	color string
	size  float32
}

type info struct {
	FirstName, LastName, Email string
	Age                        int
	job                        string
	ptshirt                    []tshirt
}

func main() {
	data := info{
		FirstName: "Mamad",
		LastName:  "dvatgar",
		Email:     "gotcha@gmail.com",
		Age:       25,
		job:       "programer",
		ptshirt: []tshirt{
			tshirt{
				color: "blue",
				size:  34.6,
			},
		},
	}

	file, _ := json.Marshal(data)
	_ = ioutil.WriteFile("go.json", file, 0644)
}
