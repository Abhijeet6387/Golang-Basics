package main

import (
	"fmt"
	"sync"
)

// Refer to documentation for more about goroutines (handled bu goruntime), waitgroups and mutex (mutual exclusion locks)
func main() {
	fmt.Println("Goroutine, Wait groups and Mutex")

	var score = []int{0}												// slice of scores, aim is to append numbers to score with multiple go routines (like threads)
	wg := &sync.WaitGroup{}												// Wait groups from sync package - To wait for multiple goroutines to finish, we can use a wait group.
	mut := &sync.Mutex{}												// There is a race condition with multiple go routines, we can use time.Sleep() to handle but it's expensive, thus mutex is used from sync package

	wg.Add(1)															// .Add()- It increases WaitGroup counter by given integer value.
	// Lambda functions - Declared and Used on the same line
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("One go routine")
		mut.Lock()														// locking at Write operation
		score = append(score, 1)
		mut.Unlock()													// unlocking after operation is completed
		wg.Done()														//.Done()- It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
	}(wg, mut)
	
	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("Two go routine")
		mut.Lock()														// locking at Write operation
		score = append(score, 2)
		mut.Unlock()													// unlocking after operation is completed
		wg.Done()
	}(wg, mut)
	
	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex ){
		fmt.Println("Three go routine")
		mut.Lock()														// locking at Write operation
		score = append(score, 3)
		mut.Unlock()													// unlocking after operation is completed
		wg.Done()
	}(wg, mut)

	wg.Wait()															//.Wait()- It Blocks the execution until it’s internal counter becomes 0.
	
	mut.Lock()															// locking at read operation
	fmt.Println(score)	
	mut.Unlock()														// unlocking after operartion is completed
}