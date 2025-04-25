# Test depedency trick

Package `bar` is depended by package `foo`.

```bash
$ diff -c bar/bar_test.go foo/foo_test.go 
*** bar/bar_test.go	2025-04-25 10:29:33.256268617 +0300
--- foo/foo_test.go	2025-04-25 10:43:22.138244798 +0300
***************
*** 1,4 ****
! package bar_test
  
  import (
  	"fmt"
--- 1,4 ----
! package foo_test
  
  import (
  	"fmt"

$ go test ./... -v
# github.com/bozaro/test-vs-compilation/foo
foo/foo.go:7:11: v.Unreachable undefined (type bar.Bar has no field or method Unreachable)
=== RUN   TestSimple
42
--- PASS: TestSimple (0.00s)
PASS
ok  	github.com/bozaro/test-vs-compilation/bar	0.002s
FAIL	github.com/bozaro/test-vs-compilation/foo [build failed]
FAIL
```
