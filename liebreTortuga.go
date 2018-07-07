package main

import (
	"fmt"
	"math/rand"
	"time"
)

var speedL = 12

var speedT = 5

var chanLiebre chan int

var chanTortuga chan int

var end chan bool

func Liebre( /*chanLiebre chan int*/ ) {
	for {
		time.Sleep(100 * time.Millisecond)
		posL := <-chanLiebre
		num := rand.Intn(100)
		if num <= 10 {
			time.Sleep(time.Second)
		} else {
			posL += speedL
		}
		chanLiebre <- posL
	}
}

func Tortuga( /*chanTortuga chan int*/ ) {
	for {
		time.Sleep(100 * time.Millisecond)
		posT := <-chanTortuga
		posT += speedT
		chanTortuga <- posT
	}
}

func checkMorder( /*chanLiebre, chanTortuga chan int*/ ) {
	for {
		posL := <-chanLiebre
		posT := <-chanTortuga

		if posL == posT {
			fmt.Println("posL: ", posL)
			fmt.Println("posT", posT)
			fmt.Println("La Tortuga Mordio la Liebre!!")
			posL -= 20
		}
		chanLiebre <- posL
		chanTortuga <- posT
	}
}

func checkGanador( /*chanLiebre, chanTortuga chan int*/ ) {
	for {

		posL := <-chanLiebre
		posT := <-chanTortuga

		if posL >= 200 {
			fmt.Println("La Liebre Gana!!")
			end <- true
			return
		}
		if posT >= 200 {
			fmt.Println("La Tortuga Gana!!")
			end <- true
			return
		}
		chanLiebre <- posL
		chanTortuga <- posT

	}
}

func Carrera( /*chanLiebre, chanTortuga chan int*/ ) {
	fmt.Println("Empieza la Carrera!!!")
	go Liebre()
	go Tortuga()
	go checkMorder()
	go checkGanador()
}

func main() {
	chanLiebre = make(chan int, 1)
	chanTortuga = make(chan int, 1)
	end = make(chan bool)
	chanLiebre <- 0
	chanTortuga <- 0

	Carrera()
	<-end
}
