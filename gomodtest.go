package gomodtest

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	fmt.Sprintf("Hi first, %s", name)
	return fmt.Sprintf("Hi, %s", name)
}
