package debug

import (
	"testing"
	"os"
	"fmt"
	"regexp"
)

type CustomWrite struct {
	str string
}

func (c *CustomWrite) Write(p []byte) (n int, err error) {
	s := fmt.Sprintf("%s", p)
	c.str += s
	return len(s), nil
}

func (c *CustomWrite) Erase() {
	c.str = ""
}

var color = "\033\\[(\\d+)m"
var reset = "\033\\[0m"

func TestDebug_Log(t *testing.T) {
	w := &CustomWrite{""}

	d := NewDebug("test")
	d.SetWriter(w, true)

	normal(w, d, t)
	withoutColor(d, w, t)
	withoutLatency(d, w, t)
	notEnabled(w, d, t)
	withColorAndWithoutLatency(d, w, t)
	withoutColorAndWithLatency(d, w, t)
	globallyDisabled(d, w, t)
}

// mustMatch is an utility function to check tests with lot of parameters that determine the state
// str in my custom writer.
func mustMatch(write *CustomWrite, debug *Debug, rex string, messageError string, t *testing.T, msg string) {
	write.Erase()
	re := regexp.MustCompile(rex)
	debug.Log(msg)
	fmt.Print(write.str)
	if (rex == "" && rex != write.str) || !re.MatchString(write.str) {
		t.Log("result: " + write.str)
		t.Log(messageError)
		t.Fail()
	}
	write.Erase()
}

// globallyDisabled test a debug that is disabled
func globallyDisabled(d *Debug, w *CustomWrite, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	Disable()
	d.Option.Reset()
	mustMatch(
		w, d, "",
		"globallyDisabled: wtf i say NOTHING should be print !",
		t, "not printing message")
}

// withoutColorAndWithLatency test a debug without color and with latency
func withoutColorAndWithLatency(d *Debug, w *CustomWrite, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	d.Option.Reset()
	mustMatch(
		w, d,
		fmt.Sprintf("^\\[%s\\]%s %s .+%s\n$", d.Name, reset, "without color and with latency", reset),
		"withoutColorAndWithLatency: not match without color and with latency",
		t, "without color and with latency")
}

// withColorAndWithoutLatency test a debug with color and without latency
func withColorAndWithoutLatency(d *Debug, w *CustomWrite, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	os.Setenv("DEBUG_COLORS", "*")
	os.Setenv("DEBUG_HIDE_LATENCY", "*")
	d.Option.Reset()
	mustMatch(
		w, d,
		fmt.Sprintf("^%s\\[%s\\]%s %s %s%s\n$", color, d.Name, reset, "with color and without latency", color, reset),
		"withColorAndWithoutLatency: not match with color and without latency",
		t, "with color and without latency")
}

// notEnabled test a debug that is not enabled
func notEnabled(w *CustomWrite, d *Debug, t *testing.T) {
	resetEnv()
	d.Option.Reset()
	mustMatch(
		w, d, "",
		"notEnabled: seems to print something, wtf ?",
		t, "this msg will not be printed")
}

// withoutLatency test a debug without latency
func withoutLatency(d *Debug, w *CustomWrite, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	os.Setenv("DEBUG_COLORS", "*")
	os.Setenv("DEBUG_HIDE_LATENCY", "*")
	d.Option.Reset()
	mustMatch(
		w, d,
		fmt.Sprintf("^%s\\[%s\\]%s %s %s%s\n$", color, d.Name, reset, "without latency!", color, reset),
		"withoutLatency: not match without latency",
		t, "without latency!")
}

// withoutColor test a debug without color
func withoutColor(d *Debug, w *CustomWrite, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	d.Option.Reset()
	mustMatch(
		w, d,
		fmt.Sprintf("^\\[%s\\]%s %s .+%s\n$", d.Name, reset, "without color", reset),
		"withoutColor: not match without color test case",
		t, "without color")
}

// normal test a simple normal default output
func normal(w *CustomWrite, d *Debug, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	os.Setenv("DEBUG_COLORS", "*")
	d.Option.Reset()
	mustMatch(
		w, d,
		fmt.Sprintf("^%s\\[%s\\]%s %s %s.+%s\n$", color, d.Name, reset, "hello world", color, reset),
		"normal: not match simple test case",
		t, "hello world")
}

func resetEnv() {
	os.Setenv("DEBUG", "")
	os.Setenv("DEBUG_HIDE_DATE", "")
	os.Setenv("DEBUG_COLORS", "")
	os.Setenv("DEBUG_HIDE_LATENCY", "")
}