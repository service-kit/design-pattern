package filter_pattern

import (
	"fmt"
	"testing"
)

func Test_FilterPattern(t *testing.T) {
	persons := []Person{
		{"Robert", MALE, SINGLE},
		{"John", MALE, MARRIED},
		{"Laura", FEMALE, SINGLE},
		{"Diana", FEMALE, MARRIED},
		{"Mike", MALE, SINGLE},
		{"Kevin", FEMALE, SINGLE},
	}
	mf := NewMaleFilter()
	ff := NewFemaleFilter()
	sf := NewSingleFilter()
	mp := mf.Filter(persons)
	fmt.Println("male persons ", mp)
	fp := ff.Filter(persons)
	fmt.Println("female persons ", fp)
	sp := sf.Filter(persons)
	fmt.Println("single persons ", sp)
	msf := NewAndFilter(mf, sf)
	msfp := msf.Filter(persons)
	fmt.Println("male single persons ", msfp)
	mosf := NewOrFilter(mf, sf)
	mosfp := mosf.Filter(persons)
	fmt.Println("male or single persons ", mosfp)
}
