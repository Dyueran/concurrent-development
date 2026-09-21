/*
Author: Yueran DING Eileen
Help received from: Felica
Help given to: Felica
License:MIT
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

//make struct containing channel
//add init, acquire, acquire and release

func main() {
	maxGoroutines := 4
	semaphore := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < 14; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()

			fmt.Printf("Running task %d\n", num)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
