package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	table := make(chan *Ball)
	done := make(chan *Ball)

	go player("Ber", table, done)
	go player("shel", table, done)

	referree(table, done)

}

type Ball struct {
	hits       int
	lastPlayer string
}

func referree(table chan *Ball, done chan *Ball) {

	table <- new(Ball)

	for {
		select {
		case Ball := <-done:
			log.Println("winner is", Ball.lastPlayer)
			return
		}
	}
}

func player(name string, table chan *Ball, done chan *Ball) {
	for {

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case Ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				log.Println(name, "drop the ball")
				done <- Ball
				return
			}
			Ball.hits++
			Ball.lastPlayer = name
			log.Println(name, "hits the ball", Ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- Ball
		case <-time.After(2 * time.Second):
			return
		}
	}
}
