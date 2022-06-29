package main

import (
	"fmt"
	"io"
)

type dz []byte

func (dz) Write(p []byte) (n int, err error) {
	return
}
func (dz) Read(p []byte) (n int, err error) {
	return

}

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Writer
	Reader
}

func main() {

	var a dz

	io.WriteString(a, "piu piu")
	mssg, _ := io.ReadAll(a)

	fmt.Println(string(mssg))
}
