package ds

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) Push(arg int) {
	for i := 0; true; i++ {
		s.arr[i+1] = s.arr[1]

	}
	s.arr[0] = arg

}

func (s *Stack) Pop() int {
	tmp := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return tmp
}

func (s *Stack) Read(index int) int {
	return s.arr[index]
}

func (s *Stack) Print() {
	fmt.Println(s.arr)
}
