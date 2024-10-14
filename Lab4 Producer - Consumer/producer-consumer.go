// Author: Gábor Major (c00271548@setu.ie)
// Date creation: 2024-10-14
// Description:
// Code for showing the Producer Consumer Problem

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Producers create a random amount of random integers
func producer(workChan chan int, producerWG *sync.WaitGroup) {
	time.Sleep(time.Second * time.Duration(rand.IntN(5)+1)) // Wait for 1 to 5 seconds
	for range rand.IntN(10) + 1 {                           // Run 1 to 10 times
		workChan <- rand.IntN(10) + 1 // Add 1 to 10 into channel
	}
	producerWG.Done()
}

// Consumers keep running until Channel is closed
func consumer(workChan chan int, consumerWG *sync.WaitGroup) {
	time.Sleep(time.Second * time.Duration(rand.IntN(5)+1)) // Wait for 1 to 5 seconds
	for {
		value, open := <-workChan // Receive value and boolean saying if channel is open
		// The open boolean only returns false if the channel is closed and empty
		if !open {
			// If channel is closed, then signal to WaitGroup and finish
			consumerWG.Done()
			return
		}
		fmt.Println(value)
	}
}

func main() {
	// Declare variables
	numberOfProducers := 5
	numberOfConsumers := 3

	workChan := make(chan int, 10)

	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup
	producerWG.Add(numberOfProducers)
	consumerWG.Add(numberOfConsumers)

	// Create producers and consumers
	for range numberOfProducers {
		go producer(workChan, &producerWG)
	}
	for range numberOfConsumers {
		go consumer(workChan, &consumerWG)
	}

	// Wait for all producers
	producerWG.Wait()
	fmt.Println("Producers done.")
	// Channel is closed
	close(workChan)
	fmt.Println("Channel closed.")
	// Wait for all consumers
	consumerWG.Wait()
	fmt.Println("Consumers done.")
}
