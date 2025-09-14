package main

// import (
// 	"fmt"
// )

// func main() {
// 	var wc WriterCloser = myWriterCloser{}
// 	fmt.Println(wc)
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type Closer interface {
// 	Close() error
// }

// type WriterCloser interface {
// 	Writer
// 	Closer
// }

// type myWriterCloser struct{}

// func (mwc myWriterCloser) Write(data []byte) (int, error) {
// 	return 0, nil
// }

// func (mwc myWriterCloser) Close() error {
// 	return nil
// }

// basics
// interfaces are a type just like structs or type aliases. structs are data containers, interfaces describe behaviours - this means we store method definitions in interfaces
// type Writer interface {
// 	Write([]byte) (int, error)                     this is a method "write" which accepts a slice of byte and then returns an integer and an error. anything that implements this interface is going to take in that slice of bytes, write it to maybe the console. TCP connection, file system etc. and then an integer comes back which is the number of bytes written, the error is there incase something goes wrong with the Write operation.
// }

// type ConsoleWriter struct {                     in GO we dont explicitly implement our interfaces, we implicitly implement it. there is no implement keyword. we create a method on our ConsoleWriter that contains the signature of our Writer interface.

// func (cw ConsoleWriter) Write(data []byte) (int, error) {             weve got a method called Write, its accepting a slice of bytes, its returning an integer and an error.
//	   n, err := fmt.Println(string(data))                               here we are juct cnverting that byte slice into a string and then print it out to the console
//	   return n,err
//}
// }

// func main() {
// 	var w Writer = ConsoleWriter{}                    the "w" variable is holding a writer which is something that implements the writer interface. at this point we dont know the type.
// 	w.Write([]byte("Hello Go!"))                      we domt also know what is written in the interface. the ConsoleWriter can be replaced by TCPWriter or whatever i want. anything that wants to use the w variable knows we can write to it.
// }

// naming conversions
// the name of the interface shows what the interface does for you.
// but if youve got single method interfaces, then the convention is to name the interface, the method name +er. lets say we are defining a write method, then the intefrface name should be writer. if we are creating an interface with a read method on it, then the interface name should be a reader
// we dont need to use structs everytime, any type that can have a method associated with it can be used, and any type can have methods associated with it

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myInt := IntCounter(0)
// 	var inc Incrementer = &myInt
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(inc.Increment())
// 	}
// }

// type Incrementer interface {
// 	Increment() int
// }

// type IntCounter int

// func (ic *IntCounter) Increment() int {
// 	*ic++
// 	return int(*ic)
// }

//package main

// import (
// 	"bytes"
// 	"fmt"
// )

// composing interfaces: how to compose interfaces together

// func main() {
// 	var wc WriterCloser = NewBufferedWriterCloser()
// 	wc.Write([]byte("Hello Youtube Listeners, this is a test"))
// 	wc.Close()
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type Closer interface {
// 	Close() error
// }

// type WriterCloser interface {
// 	Writer
// 	Closer
// }

// type BufferedWriterCloser struct {
// 	buffer *bytes.Buffer
// }

// func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
// 	n, err := bwc.buffer.Write(data)
// 	if err != nil {
// 		return 0, err
// 	}

// 	v := make([]byte, 8)
// 	for bwc.buffer.Len() > 8 {
// 		_, err := bwc.buffer.Read(v)
// 		if err != nil {
// 			return 0, err
// 		}
// 		_, err = fmt.Println(string(v))
// 		if err != nil {
// 			return 0, err
// 		}
// 	}

// 	return n, nil
// }

// func (bwc *BufferedWriterCloser) Close() error {
// 	for bwc.buffer.Len() > 0 {
// 		data := bwc.buffer.Next(8)
// 		_, err := fmt.Println(string(data))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func NewBufferedWriterCloser() *BufferedWriterCloser {
// 	return &BufferedWriterCloser{
// 		buffer: bytes.NewBuffer([]byte{}),
// 	}
// }

// type conversions
// func main() {
// 	var wc WriterCloser = NewBufferedWriterCloser()                       creating the variable
// 	wc.Write([]byte("Hello Youtube Listeners, this is a test"))           writing a couple of string in it
// 	wc.Close()                                                            implementing the close method on it.

// 	bwc := wc.(*BufferedWriterCloser)                                     here is the type conversion, we have an object, then a dot, then in the parenthesis we have a type that we are going to try and convert it to.
// 	fmt.Println(bwc)                                                      since we need to work with every variable declared, we print it here and it prints out the address of the bufferedwritercloser. so now i can work with bufferedwritercloser and not only writercloser
// }

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// )

// func main() {
// 	var wc WriterCloser = NewBufferedWriterCloser()
// 	wc.Write([]byte("Hello Youtube Listeners, this is a test"))
// 	wc.Close()

// 	bwc := wc.(io.Reader)                            this makes the application to panic, because this is another interface, this io.read interface requires a read method on it. the panic is an interface conversion, it cannot figure out how to cast the BufferedWriterCloser into the io.Reader
// 	fmt.Println(bwc)
// }

// how do we get around it (how do we convert interfaces)
// func main() {
// 	var wc WriterCloser = NewBufferedWriterCloser()
// 	wc.Write([]byte("Hello Youtube Listeners, this is a test"))
// 	wc.Close()

// 	r, ok := wc.(io.Reader)                                          here we rename the interface, we add a , ok syntax , this gives a boolean result, and then we can test against that to see if we can work with it,
// 	if ok {                                                          if conversion succeeds, we get an ok value back oout
// 		fmt.Println(r)
// 	} else {                                                         if converwsion fails, we get the zero value of whatever type we were trying to convert to. io.reader is an interface, so its zero value is nil
// 		fmt.Println("Conversion failed")
// 	}
// }

// the empty interface
// this is an interface that doesnt have any method on it.
// example:
// func main() {
// 	var myobj interface{} = NewBufferedWriterCloser()                             the syntax interface{} describes it. it is created by just deleting all the methods on it. nothing special about it, just an interface defined on the fly that has no methods on it. the nice thing is that everything can be cast on it even primitives. this is very useful when you have many things that you want to be working with but they arent type compatible with one another.
// 	if wc, ok := myobj.(WriterCloser); ok {                                       in other to do anything with a variable that has the type of an empty interface, you will have to do either type conversion or using the reflect package in other to figure out what kind of an object you are dealing with.
// 		wc.Write([]byte("Hello Youtube Listeners, this is a test"))
// 		wc.Close()
// 	}

// 	r, ok := myobj.(io.Reader)
// 	if ok {
// 		fmt.Println(r)
// 	} else {
// 		fmt.Println("Conversion failed")
// 	}
// }
// an empty interface will almost always be an intermediate stack, you declare a variable with an empty interface in it and then you find out what you recieve.
// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// )

// func main() {
// 	var myobj interface{} = NewBufferedWriterCloser()
// 	if wc, ok := myobj.(WriterCloser); ok {
// 		wc.Write([]byte("Hello Youtube Listeners, this is a test"))
// 		wc.Close()
// 	}

// 	r, ok := myobj.(io.Reader)
// 	if ok {
// 		fmt.Println(r)
// 	} else {
// 		fmt.Println("Conversion failed")
// 	}
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type Closer interface {
// 	Close() error
// }

// type WriterCloser interface {
// 	Writer
// 	Closer
// }

// type BufferedWriterCloser struct {
// 	buffer *bytes.Buffer
// }

// func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
// 	n, err := bwc.buffer.Write(data)
// 	if err != nil {
// 		return 0, err
// 	}

// 	v := make([]byte, 8)
// 	for bwc.buffer.Len() > 8 {
// 		_, err := bwc.buffer.Read(v)
// 		if err != nil {
// 			return 0, err
// 		}
// 		_, err = fmt.Println(string(v))
// 		if err != nil {
// 			return 0, err
// 		}
// 	}

// 	return n, nil
// }

// func (bwc *BufferedWriterCloser) Close() error {
// 	for bwc.buffer.Len() > 0 {
// 		data := bwc.buffer.Next(8)
// 		_, err := fmt.Println(string(data))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func NewBufferedWriterCloser() *BufferedWriterCloser {
// 	return &BufferedWriterCloser{
// 		buffer: bytes.NewBuffer([]byte{}),
// 	}
// }

// type switches
// func main() {
// 	var i interface{} = "tee"                         here we have "i" defined as an empty interface, set to the integer 0
// 	switch i.(type) {                                 this is the switch block used and its syntax, this is called a type switch, each of the cases in the switch are going to be data types.
// 	case int:
// 		fmt.Println("i is an integer")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("i dont know what i is")
// 	}
// }
// this type switches are commonly paired with empty interface in order to list out the types you are expecting to recieve, and then you will add the logic of how to get it

// why we can convert our WriterCloser into a pointer of our bufferedWriterCloser and not to the value itself
// when we define types and we assign methods to them, each one of those types has a method set. when working with types directly, the method set is all of the methods regardless of the reciever type associated with that type.
// when you implement interfaces with the concrete value directly, the variable will be defined as hokding the value my writercloser
// N/B WHEN IMPLEMENTING AN INTERFACE, if i use a value type the methods that implement the interface have to all have value recievers,
// if implementing the interface with a pointer, then you just have to have the methods there regardless of the reciever type
// so the method set for a value type is the set of all methods that have a value reciever.
// but the method for a pointer type is all of the methods with value recievers as well as all of the methods with pointer recievers
// implementing with values and pointers.

// Best Practices
// use many small interfaces, instead of large monolytic ones
// example of single method interfaces that are powerful and flexible, io.Writer, io.Reader, interface{}
// io.reader => 1 method
// io.writer => 1 method
// empty interface => 0 method                    these may be the most powerful interface in the entire language

// if you don need to export interfaces yourself, if there is no particular reason to do it, dont export it
// however export interfaces for types that you will be using in the package. if you are going to pull a value in, go ahead and accept an interface instead of a concrete value type, export interfaces for types tat will be used by package. in GO interfaces are implicitly implemented
// design functions and methods to recieve interfaces whenever possible. when accepting behaviour providers it is better to accept them as interfaces
//
