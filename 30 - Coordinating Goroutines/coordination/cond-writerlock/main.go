package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}
var readyCond = sync.NewCond(&rwmutex)

var squares = map[int]int{}

func generateSquares(max int) {
	Printfln("Generating data...")

	rwmutex.Lock()
	for i := 0; i < max; i++ {
		squares[i] = int(math.Pow(float64(i), 2))
	}
	rwmutex.Unlock()

	Printfln("Broadcasting...")
	readyCond.Broadcast()

	waitGroup.Done()
}

func readSquares(id, max, iterations int) {
	Printfln("Read data...")
	readyCond.L.Lock()
	for len(squares) == 0 {
		Printfln("wait...")
		readyCond.Wait()
	}
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()

	waitGroup.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}

	waitGroup.Add(1)
	go generateSquares(10)

	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
