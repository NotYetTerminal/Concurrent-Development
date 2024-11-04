//Barrier.go Template Code
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
// Last modify on: 2024-09-30
// Description:
// A simple barrier implemented using mutex and semaphore
// Issues: None
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, sem *semaphore.Weighted, wg2 *sync.WaitGroup, totalRoutines int64) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	//we wait here until everyone has completed part A
	sem.Release(1)
	if sem.TryAcquire(totalRoutines) { // once all the routines have finished last one comes here
		wg2.Done()                 // last routine unlocks the 2nd waitgroup
		sem.Release(totalRoutines) // reset semaphore
	} else {
		wg2.Wait() // wait until last routine
	}

	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	//we will need some of these
	ctx := context.TODO()
	sem := semaphore.NewWeighted(int64(totalRoutines))

	runAmount := 5
	for i := range runAmount {
		wg.Add(totalRoutines)
		wg2.Add(1)

		err := sem.Acquire(ctx, 10)
		if err != nil {
			fmt.Println(err)
			return
		}
		for j := range totalRoutines { //create the go Routines here
			go doStuff(j, &wg, sem, &wg2, int64(totalRoutines))
		}

		wg.Wait() //wait for everyone to finish before exiting
		fmt.Println("Loop", i+1, "is finished.")
	}
}
