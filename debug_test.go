package tests

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

	d.SetWriter(w, false)
	Enable()
	notATty(d, w, t)
	notATtyWithoutDate(d, w, t)
}

func TestRegister(t *testing.T) {
	Register("hi")
	_, err := Register("hi")
	if err == nil {
		t.Log("error on creating a debug, should ne already exist")
		t.Fail()
	}
	fmt.Printf("Hi already exist and error is %s\n", *err)
}

func TestDebug_SetFdWriter(t *testing.T) {
	d := NewDebug("Hola")
	d.SetFdWriter(os.Stdout)
}

func TestGet(t *testing.T) {
	Register("hello")

	if  _, err := Get("hello"); err != nil {
		t.Log("error on creating a debug and getting it")
		t.Fail()
	}
	if _, err := Get("helloworld"); err == nil {
		t.Log("error on getting a non registered debug")
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	Register("lol")
	err := Delete("lol")

	if err != nil {
		t.Log("error on deleting a debug")
		t.Fail()
	}
	err = Delete("loli")
	if err == nil {
		t.Log("error on deleting a debug non registered")
		t.Fail()
	}
}

func TestNewOptionDebug(t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*,-test")
	if NewOptionDebug("test").Enabled {
		t.Fail()
	}
}

// mustMatch is an utility function to check tests with lot of parameters that determine the state
// str in my custom writer.
func mustMatch(write *CustomWrite, debug *Debug, rex string, messageError string, t *testing.T, msg string) {
	write.Erase()
	re := regexp.MustCompile(rex)
	debug.Log(msg)
	if (rex == "" && rex != write.str) || !re.MatchString(write.str) {
		t.Log("result: " + write.str)
		t.Log(messageError)
		t.Fail()
	}
	write.Erase()
}

// notATty test a debug when output is not a tty
func notATtyWithoutDate(d *Debug, w *CustomWrite, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	os.Setenv("DEBUG_HIDE_DATE", "*")
	d.Option.Reset()
	mustMatch(
		w, d,
		fmt.Sprintf("^\\[%s\\] %s\n$", d.Name, "i print something to file normally"),
		"notATty: should print something without color, without latency and without date",
		t, "i print something to file normally")
}

// notATty test a debug when output is not a tty
func notATty(d *Debug, w *CustomWrite, t *testing.T) {
	resetEnv()
	os.Setenv("DEBUG", "*")
	d.Option.Reset()
	mustMatch(
		w, d,
		fmt.Sprintf("^.+\\[%s\\] %s\n$", d.Name, "i print something to file normally"),
		"notATty: should print something without color and without latency",
		t, "i print something to file normally")
}

// globallyDisabled test a debug when debug are disabled
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

// notEnabled test a debug that is not globalEnabled
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