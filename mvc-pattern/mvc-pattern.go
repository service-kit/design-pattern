package mvc_pattern

import "fmt"

type Student struct {
	Name string
	Age  int
}

type StudentView struct {
	stu *Student
}

func (s *StudentView) Show() {
	fmt.Printf("Name:%v Age:%v\n", s.stu.Name, s.stu.Age)
}

type StudentController struct {
	stu *Student
}

func (s *StudentController) AddAge() {
	s.stu.Age++
}
