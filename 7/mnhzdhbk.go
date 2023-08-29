package seven

import "fmt"

func seven() {

	for i := 20; 1 < 40; i++ {
		if i%2 == 0 {
			continue
		}
		if i%9 == 00 {
			break
		}
		fmt.Println(i)
	}
}
