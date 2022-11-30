package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}

func doSum(count int, val *int) {
	time.Sleep(time.Millisecond * 500)
	for i := 0; i < count; i++ {
		mutex.Lock()
		*val++
		mutex.Unlock()
	}
	waitGroup.Done()
}

func main() {
	counter := 0

	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter)
	}
	waitGroup.Wait()

	Printfln("Total: %v", counter)
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
