// Author: Gábor Major (c00271548@setu.ie)
// Date creation: 2024-10-21
// Description:
// Code for solving the dining philosophers problem,
// using a control channel

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const numberOfPhilosophers int = 5

// Philosopher thinks for random amount of time
func think(id int) {
	fmt.Println("Philosopher:", id, "is thinking")
	time.Sleep(time.Second * time.Duration(rand.IntN(5)+1)) // Wait for 1 to 5 seconds
}

// Philosopher eats for random amount of time
func eat(id int) {
	fmt.Println("Philosopher:", id, "is eating")
	time.Sleep(time.Second * time.Duration(rand.IntN(5)+1)) // Wait for 1 to 5 seconds
}

// Return left and right fork indexes
func left(id int) int  { return id }
func right(id int) int { return (id + 1) % numberOfPhilosophers }

// Philosopher gets its left and right fork
func getForks(id int, semChannel chan bool, forks map[int]chan bool) {
	semChannel <- true
	forks[left(id)] <- true
	forks[right(id)] <- true
}

// Philosopher puts down its left and right fork
func putForks(id int, semChannel chan bool, forks map[int]chan bool) {
	<-forks[left(id)]
	<-forks[right(id)]
	<-semChannel
}

// Philosopher runs loop and then signals it is done
func philosopher(id int, semChannel chan bool, forks map[int]chan bool) {
	for { // Run forever
		think(id)
		getForks(id, semChannel, forks)
		eat(id)
		putForks(id, semChannel, forks)
	}
}

func main() {
	// Create control channel, smaller than total number
	semChannel := make(chan bool, numberOfPhilosophers-1)
	// Create fork for each philosopher
	forks := make(map[int]chan bool)
	for i := range numberOfPhilosophers {
		forks[i] = make(chan bool, 1)
	}

	// Create philosophers
	for i := range numberOfPhilosophers {
		go philosopher(i, semChannel, forks)
	}

	// WaitGroup to allow goroutines in the background for infinity
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
