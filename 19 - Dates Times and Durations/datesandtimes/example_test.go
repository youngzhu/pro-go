package main

import (
	"fmt"
	"time"
)

func printTime(label string, t *time.Time) {
	Printfln("%s: Day: %v, Month: %v Year:%v",
		label, t.Day(), t.Month(), t.Year())
}

func Example_helloTime() {
	now := time.Now()
	specific := time.Date(2001, time.January, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(143228090, 0)

	printTime("Now", &now)
	printTime("Specific", &specific)
	printTime("UNIX", &unix)

	// Output:
	//Now: Day: 23, Month: November Year:2022
	//Specific: Day: 9, Month: January Year:2001
	//UNIX: Day: 17, Month: July Year:1974
}

func formatTime(label string, t *time.Time) {
	layout := "Day: 02, Month: Jan, Year: 2006"
	fmt.Println(label, t.Format(layout))
}
func Example_format() {
	now := time.Now()
	specific := time.Date(2001, time.January, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(143228090, 0)

	formatTime("Now:", &now)
	formatTime("Specific:", &specific)
	formatTime("UNIX:", &unix)

	// Output:
	//Now: Day: 23, Month: Nov, Year: 2022
	//Specific: Day: 09, Month: Jan, Year: 2001
	//UNIX: Day: 17, Month: Jul, Year: 1974
}

func Example_ParseDuration() {
	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hour: %v", d.Hours())
		Printfln("Minute: %v", d.Minutes())
		Printfln("Second: %v", d.Seconds())
		Printfln("Millisecond: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}

	// output:
	//Hour: 1.5
	//Minute: 90
	//Second: 5400
	//Millisecond: 5400000
}

func writeToChannelWithSleep(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name

		// 不知道是不是example test 机制的问题，都是一次性输出，没有停顿
		time.Sleep(time.Second * 5)
	}
	// 必须关闭，否则for-range会一直等待
	close(channel)
}
func Example_gettingGoroutineSleep() {
	nameChan := make(chan string)
	go writeToChannelWithSleep(nameChan)

	for name := range nameChan {
		Printfln("Read name: %v", name)
	}

	// Output:
	//
}

// AfterFunc 只接受 没有入参没有返回的函数
func Example_AfterFunc() {
	nameChan := make(chan string)
	time.AfterFunc(time.Second*5, func() {
		writeToChannelWithoutSleep(nameChan)
	})
	for name := range nameChan {
		Printfln("Read name: %v", name)
	}

	// Output:
	//
}
func writeToChannelWithoutSleep(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
	}
	// 必须关闭，否则for-range会一直等待
	close(channel)
}
