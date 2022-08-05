package main

import "fmt"

func diff(x, y uint) uint {
	d := int(x) - int(y)
	if d < 0 {
		return -uint(d)
	}
	return uint(d)
}

func generateSwitch(from, to uint) string {
	// right arrow
	keyCode := 124
	if to < from {
		// left arrow
		keyCode = 123
	}

	s := `tell application "System Events"`
	s += "\n"

	for i := uint(0); i < diff(from, to); i++ {
		s += fmt.Sprintf("  key code %d using (control down)\n", keyCode)
		s += "  delay 0.1\n"
	}

	s += "end tell\n"

	return s
}
