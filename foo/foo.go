package foo

import "github.com/bozaro/test-vs-compilation/bar"

func BrokenFunction() int {
	var v bar.Bar
	return v.Unreachable()
}
