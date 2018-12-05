package memento_pattern

type Memento struct {
	State string
}

func (m *Memento) Memento(state string) {
	m.State = state
}

func (m *Memento) GetState() string {
	return m.State
}

type Originator struct {
	State string
}

func (o *Originator) SetState(state string) {
	o.State = state
}

func (o *Originator) GetState() string {
	return o.State
}

func (o *Originator) SaveStateToMemento() *Memento {
	return &Memento{State: o.State}
}

func (o *Originator) GetStateFromMemento(m *Memento) {
	o.State = m.GetState()
}

type CareTaker struct {
	mementos []*Memento
}

func (c *CareTaker) Add(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *CareTaker) Get(i int) *Memento {
	if i >= len(c.mementos) || i < 0 {
		return nil
	}
	return c.mementos[i]
}
