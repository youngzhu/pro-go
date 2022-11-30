package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var once = sync.Once{}

var squares = map[int]int{}

func generateSquares(max int) {
	Printfln("Generating data...")

	for i := 0; i < max; i++ {
		squares[i] = int(math.Pow(float64(i), 2))
	}
}

func readSquares(id, max, iterations int) {
	Printfln("Read data...")
	once.Do(func() {
		generateSquares(max)
	})

	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}

	waitGroup.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}

	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
