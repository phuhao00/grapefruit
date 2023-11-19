package testdata

import (
	"fmt"
	"testing"
)

type Subject struct {
	observers []Observer
	state     int
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.NotifyAllObservers()
}

func (s *Subject) NotifyAllObservers() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

type Observer interface {
	Update(int)
}

type BinaryObserver struct{}

func (b *BinaryObserver) Update(state int) {
	fmt.Printf("Binary: %b\n", state)
}

type OctalObserver struct{}

func (o *OctalObserver) Update(state int) {
	fmt.Printf("Octal: %o\n", state)
}

func TestObserver(t *testing.T) {
	s := Subject{}
	s.Attach(&BinaryObserver{})
	s.Attach(&OctalObserver{})
	s.SetState(2)

}
