# Diff

Package `diff` provides a facility to compare two strings.

The output is primitive, but a little helpful.

## Install

To install, use `go get`:

```bash
$ go get -u github.com/elpinal/diff
```

-u flag stands for "update".

## Examples

```go
package main

import (
	"fmt"

	"github.com/elpinal/diff"
)

func main() {
	fmt.Println(diff.Diff("going to go", "going with go"))
        // Output: len[11 and 13]: at 6, "t" and "w"
}
```

## Contribution

1. Fork ([https://github.com/elpinal/diff/fork](https://github.com/elpinal/diff/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[elpinal](https://github.com/elpinal)
