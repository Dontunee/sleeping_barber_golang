package src

import "log"

type Customer struct {
	name string
}

func (c *Customer) Init(name string) {
	c.name = name
}

func (c *Customer) FinishCut() {
	log.Println(c.name, "has finished his/her hair cut")
	numOfHairCuts++
}

func (c *Customer) Leave() {
	log.Println(c.name, "left the shop")
}
