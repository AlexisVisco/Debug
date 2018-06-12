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
	Option		*OptionDebug
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
		NewOptionDebug(name),
		os.Stderr,
		isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())}
}

// Log print if debug is active the message with the name of the debug and the latency between
// the last call if it was activated.
func (d *Debug) Log(message string) {
	if enabled && d.Option.Enabled {
		d.writer.Write([]byte(d.Sprint(message)))
	}
	t := time.Now()
	d.LastCall = &t
}

// Sprint return the full string that should be printed.
func (d *Debug) Sprint(message string) string {
	if d.tty {
		return fmt.Sprintf("%s[%s]\033[0m %s %s%s\033[0m\n", d.color(), d.Name, message, d.color(), d.since())
	}
	return fmt.Sprintf("%s [%s] %s\n", d.date(), d.Name, message)
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
	if d.Option.Latency {
		if d.LastCall == nil {
			str = "0.00µs"
		} else {
			str = time.Since(*d.LastCall).Round(time.Millisecond).String()
		}
		return str
	}
	return ""
}

// date return the date if output is not a terminal and date is not disable.
func (d *Debug) date() interface{} {
	if d.Option.Date && !d.tty {
		return time.Now().String()
	}
	return ""
}

// color return the color if color was enabled
func (d *Debug) color() string {
	if d.Option.Color {
		return d.Color
	}
	return ""
}
