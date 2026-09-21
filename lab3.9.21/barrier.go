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
// Modified by:
// Issues:
// The barrier is not implemented!
//--------------------------------------------
/*
Author: Yueran DING Eileen
Help received from: Felica
Help given to: Felica
License:MIT
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var count int
var totalRoutines int = 10
var theLock sync.Mutex
var sem *semaphore.Weighted
var ctx context.Context

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	theLock.Lock()
	count = count + 1
	if count == totalRoutines {
		sem.Release(int64(totalRoutines - 1))
	}
	theLock.Unlock()

	if count != totalRoutines {
		sem.Acquire(ctx, 1)
	}
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	ctx = context.TODO()
	sem = semaphore.NewWeighted(int64(totalRoutines))
	sem.Acquire(ctx, int64(totalRoutines))
	//start 10 goroutine
	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &wg)
	}

	wg.Wait()
}
