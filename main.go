package main

import "github.com/Dontunee/sleeping_barber_golang/src"

/*Summary
Barbershop with one barber,
One Barber Chair,
1 waiting room with N chairs(n may be 0)
for waiting customers.
Rules
1) if there are no customers , barber falls asleep on the chair
2)if customer arrives while barber is working and waiting chair is occupied, customer leaves
3) if customer arrives while barber is working and waiting chair is available ,customer sits
4) If barber is done with a haircut , he checks the waiting room to see if there is customer to call in
5) if there is no customer in waiting room barber sleeps

*/

func main() {
	s := &src.Shop{}
	s.Init(1, []string{"tunde"})

	c := &src.Customer{}
	c.Init("tunde")

	go s.CutHair(c)
}
