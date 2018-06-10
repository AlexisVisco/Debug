package debug

import (
	"time"
	"os"
	"io"
	"fmt"
	"math/rand"
)

type Err *DebugError

type Debug struct {
	Color       string
	LastCall    *time.Time
	Name        string
	ShowLatency bool
	Writer      io.Writer
}

var debugs = make(map[string]*Debug)
var enabled = true

func NewDebug(name string) *Debug {
	return &Debug{
		randomColor(),
		nil,
		name,
		false,
		os.Stderr}
}

func Register(name string) (*Debug, Err) {
	deb := NewDebug(name)
	if _, err := Get(name); err == nil {
		debugs[name] = deb
		return deb, nil
	}
	return nil, DebugExist
}

func Get(name string) (*Debug, Err) {
	val, err := debugs[name]
	if err {
		return nil, DebugNotFound
	}
	return val, nil
}

func Delete(name string) Err {
	_, err := debugs[name]
	if err {
		return DebugNotFound
	}
	delete(debugs, name)
	return nil
}

func (d *Debug) Print(message string) {
	if enabled {
		d.Writer.Write([]byte(fmt.Sprintf("%s%s\033[0m %s %s%s\033[0m\n", d.Color, d.Name, message, d.Color, d.since())))
	}
}

func (d *Debug) Sprint(message string) string {
	return fmt.Sprintf("%s%s\033[0m %s %s%s\033[0m\n", d.Color, d.Name, message, d.Color, d.since())
}

func Enable() {
	enabled = true
}

func Disable() {
	enabled = false
}

func (d *Debug) since() string {
	var str string
	if d.ShowLatency {
		if d.LastCall == nil {
			str = "0.00µs"
		} else {
			str = time.Since(*d.LastCall).String()
		}
		t := time.Now()
		d.LastCall = &t
		return str
	}
	return ""
}

func randomColor() string {
	rand.Seed(time.Now().Unix())
	base := 30
	if rand.Int()%2 == 0 {
		base = 90
	}
	return fmt.Sprintf("\033[%dm", base+(rand.Intn(7-1)+1))
}
