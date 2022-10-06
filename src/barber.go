package src

import (
	"log"
	"time"
)

type Barber struct {
	name    string
	cutting chan *Customer
}

func (b *Barber) Init(name string) {
	b.name = name
	b.cutting = make(chan *Customer, 1)
}

func (b *Barber) CutHair(c *Customer) bool {
	select {
	case b.cutting <- c:
		return true
	default:
		return false
	}
}

func (b *Barber) Work(w *WaitingRoom) {
	for {
		c, ok := w.FindCustomer()
		if !ok {
			log.Println(b.name, "is napping")
			b.cutting <- sleepyTime
		} else {
			log.Println(b.name, "is cutting", c.name)
			b.cutting <- c
		}

		time.Sleep(time.Duration(time.Second * 5))

		//end cutting
		<-b.cutting

		if ok {
			c.FinishCut()
		}
	}
}
