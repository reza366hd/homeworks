package eighteen

import (
	"fmt"
)

func eighteen() {
	var pool int
	fmt.Println("enter your money")
	fmt.Scanln(&pool)

	fk := 2 * 5000
	tk := 4 * 2000
	ok := 1 * 1000
	fh := 10 * 500

	dakhil := money{}

	for dakhil.fivek = 0; true; dakhil.fivek++ {

		if pool < 5000 || fk == 0 {
			break
		} else {
			fk -= 5000
		}

	}

	for dakhil.twok = 0; true; dakhil.twok++ {
		if pool < 2000 || tk == 0 {
			break
		} else {
			tk -= 2000
		}
	}

	for dakhil.onek = 0; true; dakhil.onek++ {
		if pool < 1000 || ok == 0 {
			break
		} else {
			ok -= 1000
		}
	}

	for dakhil.fiveh = 0; true; dakhil.fiveh++ {
		if pool < 500 || fh == 0 {
			break
		} else {
			fh -= 500
		}
	}

	fmt.Print(
		dakhil.fivek, "fivek \n",
		dakhil.twok, "twok \n",
		dakhil.onek, "onek \n",
		dakhil.fiveh, "fiveh",
	)
}

type money struct {
	fiveh int
	twok  int
	onek  int
	fivek int
}
