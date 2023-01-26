// Using the writer interface (big part of the std lib)

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "hello world")
	io.WriteString(os.Stdout, "hello world")
}
