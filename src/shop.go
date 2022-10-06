package src

import "log"

var numOfHairCuts int
var sleepyTime *Customer

type Shop struct {
	barbers     []*Barber
	waitingRoom *WaitingRoom
}

func (s *Shop) Init(numOfChairs int, barbers []string) {
	s.barbers = make([]*Barber, 0, len(barbers))

	w := WaitingRoom{}
	w.Init(numOfChairs)
	s.waitingRoom = &w

	for i := 0; i < len(barbers); i++ {
		b := Barber{}
		b.Init(barbers[i])

		go b.Work(s.waitingRoom)
	}
}

func (s *Shop) CutHair(c *Customer) {
	for _, b := range s.barbers {
		log.Println("checking the barbers")
		if b.CutHair(c) {
			log.Println("customer could cut hair")
			return
		}
	}

	if s.waitingRoom.FindSeat(c) {
		log.Println("got a seat to wait")
		return
	}

	c.Leave()
}
