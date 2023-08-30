package eleven

import "fmt"

func eleven() {
	var s string
	var student []string

	for {
		fmt.Println("do you have any student? (inter no for end)")
		fmt.Scanln(&s)

		if s == "no" {
			break
		}
		student = append(student, s)

	}
	fmt.Println(student)
}
