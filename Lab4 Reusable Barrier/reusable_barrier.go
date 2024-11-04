//Lab3 Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Gábor Major (c00271548@setu.ie)
// Description:
// A reusable barrier implemented using mutex and unbuffered channels
// Issues:
// None I hope
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func doStuff(goNum int, arrived *atomic.Int64, max int, wg *sync.WaitGroup, channel1 chan bool, channel2 chan bool, theLock *sync.Mutex) bool {
	for range 3 {
		time.Sleep(time.Second)
		fmt.Println("Part A", goNum)
		//we wait here until everyone has completed part A
		arrived.Add(1)
		theLock.Lock()
		if arrived.Load() == int64(max) { //last to arrive -signal others to go
			theLock.Unlock()
			for range max - 1 { // run for all the other routines to free them up
				channel1 <- true
			}
		} else { //not all here yet we wait until signal
			theLock.Unlock()
			<-channel1
		} //end of if-else
		fmt.Println("Part B", goNum)
		// wait here until everyone has completed part B
		arrived.Add(-1)
		theLock.Lock()
		if arrived.Load() == 0 { // last routine arrives here
			theLock.Unlock()
			for range max - 1 { // run for all the other routines to free them up
				channel2 <- true
			}
		} else {
			theLock.Unlock()
			<-channel2 // wait here for last routine
		}
		fmt.Println("Part C", goNum)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 100
	var arrived atomic.Int64
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	var theLock sync.Mutex
	channel1 := make(chan bool) //use unbuffered channel in place of semaphore
	channel2 := make(chan bool)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, channel1, channel2, &theLock)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
