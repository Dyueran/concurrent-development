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
	"math/rand/v2"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Global variables shared between functions --A BAD IDEA
var count int
var theLock sync.Mutex
var sem *semaphore.Weighted
var ctx context.Context
var threadCount int = 2

func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)

	//Rendezvous here
	theLock.Lock()
	count++
	localCount := count
	if count == threadCount {
		sem.Release(1)
	}
	theLock.Unlock()

	if localCount != threadCount {
		sem.Acquire(ctx, 1)
	}

	fmt.Println("Part B", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	ctx = context.TODO()
	sem = semaphore.NewWeighted(int64(threadCount))
	sem.Acquire(ctx, int64(threadCount))

	wg.Add(threadCount)
	for N := 0; N < threadCount; N++ {
		go WorkWithRendezvous(&wg, N)
	}
	wg.Wait()
}
