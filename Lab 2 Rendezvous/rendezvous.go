// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Modified by: Gábor Major (c00271548@setu.ie)
// Date creation: 2024-11-24
// Description:
// Code for creating a rendezvous point for 2 threads.

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, chan1 chan bool, chan2 chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	if Num == 0 {
		chan1 <- true
		<-chan2
	} else if Num == 1 {
		chan2 <- true
		<-chan1
	}

	fmt.Println("Part B", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	chan1 := make(chan bool, 1)
	chan2 := make(chan bool, 1)
	threadCount := 2

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, chan1, chan2)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
}
