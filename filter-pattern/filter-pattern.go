package filter_pattern

const (
	MALE    = "male"
	FEMALE  = "female"
	SINGLE  = "single"
	MARRIED = "married"
)

type Person struct {
	Name          string
	Gender        string
	MaritalStatus string
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetGender() string {
	return p.Gender
}

func (p *Person) GetMaritalStatus() string {
	return p.MaritalStatus
}

type IFilter interface {
	Filter([]Person) []Person
	Match(Person) bool
}

type BaseFilter struct {
	MatchFunc func(Person) bool
}

func (f *BaseFilter) Match(person Person) bool {
	return f.MatchFunc(person)
}

func (f *BaseFilter) Filter(persons []Person) []Person {
	outPersons := make([]Person, len(persons))
	index := 0
	for _, p := range persons {
		if f.Match(p) {
			outPersons[index] = p
			index++
		}
	}
	return outPersons[:index]
}

func NewMaleFilter() IFilter {
	mf := new(BaseFilter)
	mf.MatchFunc = func(person Person) bool {
		return MALE == person.GetGender()
	}
	return mf
}

func NewFemaleFilter() IFilter {
	mf := new(BaseFilter)
	mf.MatchFunc = func(person Person) bool {
		return FEMALE == person.GetGender()
	}
	return mf
}

func NewSingleFilter() IFilter {
	mf := new(BaseFilter)
	mf.MatchFunc = func(person Person) bool {
		return SINGLE == person.GetMaritalStatus()
	}
	return mf
}

func NewOrFilter(f1, f2 IFilter) IFilter {
	mf := new(BaseFilter)
	mf.MatchFunc = func(person Person) bool {
		return f1.Match(person) || f2.Match(person)
	}
	return mf
}

func NewAndFilter(f1, f2 IFilter) IFilter {
	mf := new(BaseFilter)
	mf.MatchFunc = func(person Person) bool {
		return f1.Match(person) && f2.Match(person)
	}
	return mf
}
