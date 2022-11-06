/*

Race condition

A race condition situation occurs when the execution of a program is no longer deterministic,
due to the interleaving between two or more pieces of code executing concurrently.
In Go, these pieces of code are expressed using Goroutines.
It usually appears when these pieces of code share some common resource (variable, data structure, etc).

In this particular example, both Goroutines share the same string variable that is going to be
printed eventually, but it is impossible to predict the final value that it's going to take and be
printed in each execution of the program, since it’s impossible to know which Goroutine is going to be executed first and last.
So we have a non deterministic situation and a race condition due to the interleaving
between “setValue1” and “setValue2” functions executing concurrently using Goroutines.

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func setValue1(waitGroup *sync.WaitGroup, value *string) {

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	*value = "Goroutine1 wins"

	waitGroup.Done()
}

func setValue2(waitGroup *sync.WaitGroup, value *string) {

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	*value = "Goroutine2 wins"

	waitGroup.Done()
}

func main() {
	var value string
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go setValue1(&waitGroup, &value)
	go setValue2(&waitGroup, &value)

	waitGroup.Wait()
	fmt.Println(value)
}
