// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var pool int
// 	fmt.Println("enter your money")
// 	fmt.Scanln(&pool)

// 	fivek := pool / 5000
// 	pool = pool - fivek*5000
// 	twok := pool / 2000
// 	pool = pool - twok*2000
// 	onek := pool / 1000
// 	pool = pool - onek*1000
// 	fiveh := pool / 500

// 	dakhil := money{
// 		fiveh: fiveh,
// 		onek:  onek,
// 		twok:  twok,
// 		fivek: fivek,
// 	}

// 	fmt.Println(
// 		dakhil.fivek, "fivek \n",
// 		dakhil.twok, "twok \n",
// 		dakhil.onek, "onek \n",
// 		dakhil.fiveh, "fiveh \n",
// 	)
// }

// type money struct {
// 	fiveh int
// 	onek  int
// 	twok  int
// 	fivek int
// }
