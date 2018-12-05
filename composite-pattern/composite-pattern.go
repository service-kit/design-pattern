package composite_pattern

type Employee struct {
	Name         string
	Subordinates []Employee
}

func (e *Employee) Add(em Employee) {
	e.Subordinates = append(e.Subordinates, em)
}
