package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// When you write an HTTP handler, you are given an http.ResponseWriter and
// the http.Request that was used to make the request.
// When you implement your server you write your response using the writer.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

// =================== NOTES ==========================

/*
from package "fmt"
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintf(format, a)
    n, err = w.Write(p.buf)
    p.free()
    return
}

from package "io"
type Writer interface {
    Write(p []byte) (n int, err error)
}

Buffer type of the package "bytes" implements the Writer interface.
So it is a child of the Writer type, because it also implementes the Write method.

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(len(p))
	if !ok {
		m = b.grow(len(p))
	}
	return copy(b.buf[m:], p), nil
}

Therefor we implemented the following Greet method,
but this is not general purpose because we can only pass io.Writer
or the buffer child but none of the other childeren.

func Greet(writer bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
*/
