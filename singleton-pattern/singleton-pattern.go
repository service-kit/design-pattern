package singleton_pattern

import "sync"

var once sync.Once
var instance *singleton

type singleton struct {
	Value int
}

func (s *singleton) SetValue(v int) {
	s.Value = v
}

func (s *singleton) GetValue() int {
	return s.Value
}

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
