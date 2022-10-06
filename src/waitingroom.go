package src

import (
	"log"
)

type WaitingRoom struct {
	seats chan *Customer
}

func (w *WaitingRoom) Init(nc int) {
	w.seats = make(chan *Customer, nc)
}

func (w *WaitingRoom) FindSeat(c *Customer) bool {
	select {
	case w.seats <- c:
		log.Println(c.name, "sitting in the waiting room")
		return true
	default:
		log.Println(c.name, "left because the waiting room is occupied")
		return false
	}
}

func (w *WaitingRoom) FindCustomer() (*Customer, bool) {
	select {
	case c, ok := <-w.seats:
		log.Println(c.name, "is in the waiting room")
		return c, ok
	default:
		log.Println("no customer in the waiting room")
		return nil, false
	}
}
