package main

import (
	"../pkg/debug"
	"time"
	"strconv"
)

var fivesec, _ = debug.Register("five")

var five = 0


func main() {
	// manually set option of fivesec and onemin for portability
	fivesec.Option.Color = true
	fivesec.Option.Enabled = true
	fivesec.Option.Latency = true

	doEvery(5 * time.Second, func(i time.Time) {
		fivesec.Log("five = " + strconv.Itoa(five))
		five++
	})
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}