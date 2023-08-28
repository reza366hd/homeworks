package eleven

import "fmt"

func eleven() {
	var s string
	var student []string
	// Out:
	for {
		fmt.Println("do you have any student? (inter no for end)")
		fmt.Scanln(&s)

		// switch s {
		// case "no":
		// 	break Out

		// default:
		// 	student = append(student, s)
		// }

		if s == "no" {
			break
		}
		student = append(student, s)

	}
	fmt.Println(student)
}
