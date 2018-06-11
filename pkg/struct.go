package debug

import (
	"time"
	"io"
	"github.com/mattn/go-isatty"
	"os"
	"fmt"
)

type Debug struct {
	Color       string
	LastCall    *time.Time
	Name        string
	ShowLatency bool
	writer      io.Writer
	tty         bool
}

// NewDebug Create a debug structure without registering it. Cannot be accessible with [`Get`](#Get).
// Generate a random color between 31 to 37 and 91 to 97 as ANSI code.
func NewDebug(name string) *Debug {
	return &Debug{
		attributeColor(name),
		nil,
		name,
		false,
		os.Stderr,
		isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())}
}

// Log print if debug is active the message with the name of the debug and the latency between
// the last call if it was activated.
func (d *Debug) Log(message string) {
	if enabled && d.tty {
		d.writer.Write([]byte(fmt.Sprintf("%s%s\033[0m %s %s%s\033[0m\n", d.Color, d.Name, message, d.Color, d.since())))
	} else {
		d.writer.Write([]byte(fmt.Sprintf("%s\033[0m %s %s\033[0m\n", d.Name, message, d.since())))
	}
}

// Sprint return the full string that should be printed.
func (d *Debug) Sprint(message string) string {
	return fmt.Sprintf("%s%s\033[0m %s %s%s\033[0m\n", d.Color, d.Name, message, d.Color, d.since())
}

// SetWriter set the writer, if it's a terminal set to true the next parameter.
func (d *Debug) SetWriter(writer io.Writer, tty bool) *Debug {
	d.writer = writer
	d.tty = tty
	return d
}

// SetFdWriter will set the writer and determine if the file.Fd() is a terminal.
func (d *Debug) SetFdWriter(file *os.File) *Debug {
	d.writer = file
	d.tty = isatty.IsTerminal(file.Fd()) || isatty.IsCygwinTerminal(file.Fd())
	return d
}

// since return a string formatted with the time between now and the last call of since
// to be human readable.
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