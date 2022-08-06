package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	c := w.(*bytes.Buffer)
	fmt.Printf("%T", f)
	fmt.Printf("%T", c)
}
