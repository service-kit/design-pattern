package iterator_pattern

import "fmt"

type iterator interface {
	HasNext() bool
	Next() iterator
	Pre() iterator
	HasPre() bool
	SetNext(iterator)
	SetPre(iterator)
	Value() interface{}
	SetValue(interface{})
}

type Iterator struct {
	data     interface{}
	nextIter iterator
	preIter  iterator
}

func (i *Iterator) HasPre() bool {
	return nil != i.preIter
}

func (i *Iterator) HasNext() bool {
	return nil != i.nextIter
}

func (i *Iterator) SetPre(iter iterator) {
	i.preIter = iter
}

func (i *Iterator) Pre() iterator {
	return i.preIter
}

func (i *Iterator) SetNext(iter iterator) {
	i.nextIter = iter
}

func (i *Iterator) Next() iterator {
	return i.nextIter
}

func (i *Iterator) Value() interface{} {
	return i.data
}

func (i *Iterator) SetValue(data interface{}) {
	i.data = data
}

type Container interface {
	First() iterator
	End() iterator
}

type List struct {
	first iterator
	end   iterator
}

func (a *List) First() iterator {
	return a.first.Next()
}

func (a *List) End() iterator {
	return a.end
}

func (a *List) Pushback(data interface{}) {
	newIter := new(Iterator)
	newIter.data = data
	newIter.SetNext(a.end)
	a.end.Pre().SetNext(newIter)
	newIter.SetPre(a.end.Pre())
	a.end.SetPre(newIter)
	fmt.Printf("push back %v to %v\n", data, newIter.Pre().Value())
}

func (a *List) Popback() interface{} {
	if a.first.Next() == a.end {
		return nil
	}
	data := a.end.Pre().Value()
	a.end.Pre().SetNext(nil)         //1
	a.end.SetPre(a.end.Pre().Pre())  //2
	a.end.Pre().Next().SetPre(a.end) //3
	a.end.Pre().SetNext(a.end)       //4
	fmt.Printf("pop back %v\n", data)
	return data
}

func NewList() *List {
	l := new(List)
	l.first = new(Iterator)
	l.end = new(Iterator)
	l.first.SetNext(l.end)
	l.end.SetPre(l.first)
	return l
}
