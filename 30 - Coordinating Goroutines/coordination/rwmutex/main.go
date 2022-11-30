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

var squares = map[int]int{}

func calcSquares(max, iterations int) {
	for i := 0; i < iterations; i++ {
		n := rand.Intn(max)

		rwmutex.RLock()
		square, ok := squares[n]
		rwmutex.RUnlock()

		if ok {
			Printfln("Cached value: %v= %v", n, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[n]; !ok {
				squares[n] = int(math.Pow(float64(n), 2))
				Printfln("Added value: %v= %v", n, squares[n])
			}
			rwmutex.Unlock()
		}
	}
	waitGroup.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go calcSquares(10, 5)
	}
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
