package main

import "fmt"

func main() {
	c := pars{}
	a := samand{}
	d := quick{}
	defer fmt.Println("The program ended")
	sorat(a)
	sorat(d)
	sorat(c)

}

type samand struct{}
type pars struct{}
type quick struct{}

func (samand) Speed() {
	fmt.Println("samand 130k/h")
}

func (pars) Speed() {
	fmt.Println("pars 140k/h")
}

func (quick) Speed() {
	fmt.Println("speed 90k/h")
}

type ikco interface {
	Speed()
}

func sorat(model ikco) {
	model.Speed()

	j, ok := model.(pars)

	fmt.Println(j, ok)

}
