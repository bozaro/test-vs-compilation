package foo_test

import (
	"fmt"
	"github.com/bozaro/test-vs-compilation/foo"
	"testing"
)

func TestSimple(t *testing.T) {
	fmt.Println(foo.BrokenFunction())
}
