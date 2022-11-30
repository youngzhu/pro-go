package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	countKey = iota
	sleepPeriodKey
)

func processRequest(ctx context.Context, wg *sync.WaitGroup) {
	total := 0

	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)

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
			time.Sleep(sleepPeriod)
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
	ctx = context.WithValue(ctx, countKey, 10)
	ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond*200)

	go processRequest(ctx, &wg)

	wg.Wait()
	Printfln("Done.")
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
