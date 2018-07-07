package main

import (
	"fmt"
	"time"
)

var end = make(chan bool)

type Filosofo struct {
	hambre chan int
	name   string
	comer  chan bool
}

//fmt.Printf("%d: %s\n", i, f.name)
func checkHambre(fils []Filosofo) {
	for {
		max := 0
		ind := 0
		for i, f := range fils {

			select {
			//	case <-f.comer:

			default:
				//	fmt.Println(f.name)
				hambre := <-f.hambre
				//	fmt.Println(hambre)
				if hambre > max {
					max = hambre
					ind = i
				}
				f.hambre <- hambre

			}
		}
		if fils[ind].name == "Sartre" {
			//fmt.Println(fils[ind].name, "MAYOR")
		}
		fils[ind].comer <- true
		time.Sleep(time.Second)
	}
}

func eat(fil Filosofo, izq, der chan bool) {
	for {

		select {

		case <-fil.comer:
			hambre := <-fil.hambre
			hambre = 0
			fil.hambre <- hambre
			//	LEFT:
			select {
			case <-izq:
				//for {
				select {

				case <-der:
					fmt.Printf("%s esta comiendo\n", fil.name)
					time.Sleep(time.Second)
					izq <- true
					der <- true
				//	break LEFT
				default:
					izq <- true
					//	break LEFT
				}
			//	}
			default:
				//fmt.Printf("%s esta pensando\n", fil.name)
				time.Sleep(100 * time.Millisecond)

			}

		default:
			h := <-fil.hambre
			h++
			fil.hambre <- h

		}

	}
}

func dinner(fils []Filosofo, forks []chan bool) {
	for i, f := range fils {
		go eat(f, forks[i], forks[(i+1)%len(forks)])
	}
	go checkHambre(fils)
}

func main() {
	filosofos := []Filosofo{
		Filosofo{name: "Socrates"},
		Filosofo{name: "Platon"},
		Filosofo{name: "Kant"},
		Filosofo{name: "Descartes"},
		Filosofo{name: "Sartre"},
	}
	forks := make([]chan bool, len(filosofos))
	for i, _ := range forks {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}
	for i, _ := range filosofos {
		filosofos[i].hambre = make(chan int, 1)
		filosofos[i].comer = make(chan bool, 1)
		filosofos[i].hambre <- 0
	}

	dinner(filosofos, forks)

	time.Sleep(10 * time.Second)

}
