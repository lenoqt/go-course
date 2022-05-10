package main

import (
	"bytes"
	"fmt"
	"io"
)

// Implicit implementation

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Extending first example
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferredWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// With arbitrary types

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferredWriterCloser()
	wc.Write([]byte("Hello there"))
	wc.Close()

	// Type conversion
	bwc := wc.(*BufferedWriterCloser) // This has to be a pointer
	fmt.Println(bwc)
	// Avoid panic when there no available cast to a type
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	// Empty interfaces
	var myObj interface{} = NewBufferredWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("From empty interface"))
		wc.Close()
	}
	re, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(re)
	} else {
		fmt.Println("Type conversion failed")
	}

	// Type switch
	var i interface{} = 0
	switch i.(type) {
	case string:
		fmt.Println("i is an integer")
	case int:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")

	}
}
