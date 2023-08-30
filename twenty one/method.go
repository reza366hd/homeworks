package twentyone

import "fmt"

type animal struct {
	cat string
	dog string
}

func (h *animal) anm() {
	fmt.Println("hahaha", (*h).cat)
}

func twentyone() {
	anml := animal{
		cat: "meaw",
		dog: "hup",
	}

	fmt.Println(anml.anm)

}
