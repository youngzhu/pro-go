package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func processRequest(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			Printfln("Stopping process - request cancelled")
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 100)
		}
	}
	Printfln("%v Request processed", total)
end:
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel := context.WithCancel(context.Background())
	go processRequest(ctx, &wg, 10)

	time.Sleep(time.Millisecond * 500)
	Printfln("Canceling ...")
	cancel()

	wg.Wait()
	Printfln("Done.")
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
