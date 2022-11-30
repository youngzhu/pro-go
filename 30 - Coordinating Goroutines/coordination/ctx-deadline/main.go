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
			if ctx.Err() == context.Canceled {
				Printfln("Stopping process - request cancelled")
			} else {
				Printfln("Stopping process - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 200)
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

	// WithTimeout, WithDeadline 效果一样
	//ctx, _ := context.WithTimeout(context.Background(), time.Second)
	t := time.Now().Add(time.Second)
	ctx, _ := context.WithDeadline(context.Background(), t)

	go processRequest(ctx, &wg, 10)

	wg.Wait()
	Printfln("Done.")
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
