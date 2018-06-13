package tests

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
	Enabled bool	// If the Debug structure associated with this option is globalEnabled.
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
		!isIn(name, os.Getenv("DEBUG_HIDE_DATE")),
		isIn(name, os.Getenv("DEBUG_COLORS")),
		!isIn(name, os.Getenv("DEBUG_HIDE_LATENCY")),
		isIn(name, os.Getenv("DEBUG")),
	}
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
func isIn(name, val string) bool {
	hasExclude := false
	hasAMatch := false
	if val == "" {
		return false
	}
	lst := strings.Split(val, ",")
	for _, v := range lst {
		res, _ := filepath.Match(v, name)
		if !res  && strings.HasPrefix(v, "-"){
			res, _ := filepath.Match(string([]rune(v)[1:len(v)]), name)
			if res {
				hasExclude = true
			}
		} else {
			hasAMatch = true
		}
	}
	return hasAMatch && !hasExclude
}
