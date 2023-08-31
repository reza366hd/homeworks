package twentytwo

import (
	"fmt"
)

type animal struct {
	cat    string
	dog    string
	cow    string
	dock   string
	mouse  string
	fil    string
	khoros string
}

func (h *animal) caat() string {
	return (*h).cat
}

func twentytwo() {
	anml := animal{
		cat:    "meaw",
		dog:    "hup",
		cow:    "maa",
		dock:   "dock dock",
		mouse:  "ch ch",
		fil:    "hoooo",
		khoros: "gogoli gogo",
	}

	fmt.Println(anml.caat())

}
