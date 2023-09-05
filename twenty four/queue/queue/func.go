package ds

import "fmt"

type Queue struct {
	arr []int
}

func (s *Queue) Push(arg int) {

	s.arr = append([]int{arg}, s.arr...)
}

func (s *Queue) Pop() int {
	tmp := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return tmp
}

func (s *Queue) Read(index int) int {
	return s.arr[index]
}

func (s *Queue) Print() {
	fmt.Println(s.arr)
}
