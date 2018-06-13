package main

import (
	"debug"
	"time"
	"strconv"
	"sync"
)

var fivesec, _ = debug.Register("5 times")
var nivesec, _ = debug.Register("9 times")
var wait sync.WaitGroup

var five = 0
var nine = 0

func main() {
	// manually set option of fivesec for portability
	fivesec.Option.Color = true
	fivesec.Option.Enabled = true
	fivesec.Option.Latency = true
	nivesec.Option.Color = true
	nivesec.Option.Enabled = true
	nivesec.Option.Latency = true

	wait.Add(1)
	go doEvery(5 * time.Second, func(i time.Time) {
		fivesec.Log("5 = " + strconv.Itoa(five))
		five++
	})
	go doEvery(9 * time.Second, func(i time.Time) {
		nivesec.Log("9 = " + strconv.Itoa(nine))
		nine++
	})
	wait.Wait()
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}