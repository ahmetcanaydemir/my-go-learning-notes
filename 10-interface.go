/*

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}




func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}



//Buffered writer closer

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Merhaba bu çok uzun bir metin"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

 //Cast
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Dönüşüm başarısız oldu")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

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

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// Üstteki örneğin Empty Interface hali
func main() {
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Bu çok çok uzun bir metin"))
		wc.Close()
	}

	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Dönüşüm başarısız oldu")
	}
}




func main() {
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is integer")
	case string:
		fmt.Println("i string")
	default:
		fmt.Println("Bilinmiyor")
	}
}
*/

package main

import (
	"fmt"
)

func main() {
	var wc WriterCloser = myWriterCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct {
}

func (mwc myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil

}
