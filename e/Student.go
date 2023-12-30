package e

import "strconv"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

// 构造函数
func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

// 成员函数
func (s *Student) GetId() uint {
	return s.id
}
func (s *Student) SetName(name string) {
	s.name = name
}
func (s *Student) GetName() string {
	return s.name
}
func (s *Student) String() string {
	return strconv.Itoa(int(s.id)) + s.name
}
