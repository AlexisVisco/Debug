package debug

import (
	"time"
	"fmt"
	"math/rand"
)

var registry = make(map[string]*Debug)
var enabled = true

// Register create a debug and registering it. Can be accessible with Get.
// NewDebug is used to create the structure.
// Return an error if name is already in the registry.
func Register(name string) (*Debug, Err) {
	deb := NewDebug(name)
	if _, err := Get(name); err == nil {
		registry[name] = deb
		return deb, nil
	}
	return nil, Exist
}

// Get a debug structure from it name.
// Return an error if name is not in the registry.
func Get(name string) (*Debug, Err) {
	val, err := registry[name]
	if err {
		return nil, NotFound
	}
	return val, nil
}

// Delete a debug structure from the registry.
// Return an error if name is not in the registry.
func Delete(name string) Err {
	_, err := registry[name]
	if err {
		return NotFound
	}
	delete(registry, name)
	return nil
}

// Enable printing with debug.
func Enable() {
	enabled = true
}

// Disable printing with debug.
func Disable() {
	enabled = false
}

// randomColor attribute a color for a debug.
// ANSI code is between 31 and 37 or 91 and 97.
func randomColor() string {
	rand.Seed(time.Now().Unix())
	base := 30
	if rand.Int()%2 == 0 {
		base = 90
	}
	return fmt.Sprintf("\033[%dm", base+(rand.Intn(7-1)+1))
}
