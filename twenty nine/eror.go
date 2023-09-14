package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		pune(err)
	}
	if err != nil {
		pune(err)
	}
	defer f.Close()

	val := "hello"
	data := []byte(val)

	err = ioutil.WriteFile("file.txt", data, 0)
	if err != nil {
		pune(err)
	}
}

func pune(err error) {
	if err != nil {
		if errors.As(err, &A{}) {
			fmt.Println("error type A :", err.Error())
		} else if errors.As(err, &d{}) {
			fmt.Println("error type d :", err.Error())
		} else if errors.As(err, &y{}) {
			fmt.Println("error type y :", err.Error())
		} else if errors.As(err, &o{}) {
			fmt.Println("error type o :", err.Error())
		} else {
			fmt.Println("unknown error : ", err.Error())
		}
	}
}

type A struct{}

func (a A) Error() string {
	b := "file address was invalid"
	return b
}

type d struct{}

func (d d) Error() string {
	q := "invalid syntxs"
	return q
}

type y struct{}

func (y y) Error() string {
	p := "p decalred and not used"
	return p
}

type o struct{}

func (o o) Error() string {
	c := "undefind var"
	return c
}
