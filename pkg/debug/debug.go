package debug

import (
	"fmt"
)

var registry = make(map[string]*Debug)
var globalEnabled = true

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
	_, err := Get(name)
	if err != nil {
		return err
	}
	delete(registry, name)
	return nil
}

// Enable printing with debug.
func Enable() {
	globalEnabled = true
}

// Disable printing with debug.
func Disable() {
	globalEnabled = false
}

func hashJenkins(name string) uint32 {
	var hash uint32 = 0
	for e := range name {
		hash += uint32(name[e])
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return hash
}

// attributeColor attribute a color for a debug.
// ANSI code is between 31 and 37 or 91 and 97.
func attributeColor(name string) string {
	numbers := []int {31, 32, 33, 34, 35, 36, 37, 91, 92, 93, 94, 95, 96, 97}
	return fmt.Sprintf("\033[%dm", numbers[int(hashJenkins(name)) % len(numbers)])
}
