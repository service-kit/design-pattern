package observer_pattern

import "fmt"

type Observer interface {
	Update()
	Init(*Subject)
}

type Subject struct {
	state     string
	observers []Observer
}

func (s *Subject) Attach(obs Observer) {
	s.observers = append(s.observers, obs)
}

func (s *Subject) GetState() string {
	return s.state
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.notifyAllObservers()
}

func (s *Subject) notifyAllObservers() {
	for _, ob := range s.observers {
		ob.Update()
	}
}

type BinaryObserver struct {
	sub *Subject
}

func (b *BinaryObserver) Init(sub *Subject) {
	b.sub = sub
	sub.Attach(b)
}

func (b *BinaryObserver) Update() {
	fmt.Println("BinaryObserver::Update,State:", b.sub.state)
}

type OctalObserver struct {
	sub *Subject
}

func (b *OctalObserver) Init(sub *Subject) {
	b.sub = sub
	sub.Attach(b)
}

func (b *OctalObserver) Update() {
	fmt.Println("OctalObserver::Update,State:", b.sub.state)
}
