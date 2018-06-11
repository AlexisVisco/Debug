package debug

/**
NOT USED YET
NOT USED YET
NOT USED YET
NOT USED YET
NOT USED YET
NOT USED YET
 */

import (
	"os"
	"strings"
	"path/filepath"
)

// An OptionDebug represent a list of options to toggle some functionality of Debug structure.
type OptionDebug struct {
	Date    bool // Hide date from debug output (non-TTY).
	Color   bool // Whether to use colors in the debug output.
	Latency bool // Whether to display latency between last call of debug.
	Enabled bool // If the Debug structure associated with this option is enabled.
}

var env = getEnv()

// NewOptionDebug create a OptionDebug with custom parameters.
// You can use * to match all characters.
//
// DEBUG=param1,[param2,...]				Enable debug
// DEBUG_HIDE_DATE=param1,[param2,...]		Hide date on non tty output debug
// DEBUG_COLORS=param1,[param2,...]			Enable color on output
// DEBUG_HIDE_LATENCY=param1,[param2,...]	Hide latency to output
func NewOptionDebug(name string) *OptionDebug {
	opt := OptionDebug{
		!isIn(name, "DEBUG_HIDE_DATE"),
		isIn(name, "DEBUG_COLORS"),
		!isIn(name, "DEBUG_HIDE_LATENCY"),
		isIn(name, "DEBUG")}
	return &opt
}

// getEnv return a map of key=value representing the environment.
// Useful to retrieve some value for a key without time effort.
func getEnv() map[string]string {
	environement := make(map[string]string)
	for _, v := range os.Environ() {
		keys := strings.Split("=", v)
		environement[keys[0]] = keys[1]
	}
	return environement
}

// isIn transform a value of a environment key into a list of values separated with a comma.
// Then, use filepath.match to check if name is in, if -PATTERN is present, so it's a exclude.
func isIn(name, envName string) bool {
	val, err := env[envName]
	if err {
		return false
	}
	lst := strings.Split(",", val)
	for _, v := range lst {
		if res, ok := filepath.Match(v, name); ok != nil && res && !strings.HasPrefix(v, "-") {
			return true
		}
	}
	return false
}
