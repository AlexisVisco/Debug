package debug

import (
	"strings"
	"path/filepath"
	"os"
)

// An OptionDebug represent a list of options to toggle some functionality of Debug structure.
type OptionDebug struct {
	name    string	// keep name in memory for reset
	Date    bool	// Show date from debug output (non-TTY).
	Color   bool	// Whether to use colors in the debug output.
	Latency bool	// Whether to display latency between last call of debug.
	Enabled bool	// If the Debug structure associated with this option is enabled.
}

// NewOptionDebug create a OptionDebug with custom parameters.
// You can use * to match all characters.
//
// DEBUG=param1,[param2,...]				Enable debug
// DEBUG_HIDE_DATE=param1,[param2,...]		Hide date on non tty output debug
// DEBUG_COLORS=param1,[param2,...]			Enable color on output
// DEBUG_HIDE_LATENCY=param1,[param2,...]	Hide latency to output
func NewOptionDebug(name string) *OptionDebug {
	opt := OptionDebug{
		name,
		!isIn(name, "DEBUG_HIDE_DATE"),
		isIn(name, "DEBUG_COLORS"),
		!isIn(name, "DEBUG_HIDE_LATENCY"),
		isIn(name, "DEBUG")}
	return &opt
}

// Reset re-assign Enabled, Latency, Color, Date to default values.
func (opt *OptionDebug) Reset() {
	reset := NewOptionDebug(opt.name)
	opt.Enabled = reset.Enabled
	opt.Latency = reset.Latency
	opt.Color = reset.Color
	opt.Date = reset.Date
}

// isIn transform a value of a environment key into a list of values separated with a comma.
// Then, use filepath.match to check if name is in, if -PATTERN is present, it's an exclude.
func isIn(name, envName string) bool {
	val := os.Getenv(envName)
	if val == "" {
		return false
	}
	lst := strings.Split(val, ",")
	for _, v := range lst {
		res, ok := filepath.Match(v, name)
		if ok == nil && res && !strings.HasPrefix(v, "-") {
			return true
		}
	}
	return false
}
