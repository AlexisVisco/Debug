package main

import (
	debug "github.com/AlexisVisco/debug"
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