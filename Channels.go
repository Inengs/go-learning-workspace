package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // strongly typed to a struct with no fields and this requires zero memory allocation this is called a signal only channel

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} // the second curly brace is to initialize the empty struct
}

func logger() {
	for {
		select { //  what the select statement does it that the entire message is going to block until a message is recieved on one of the channels tha its listening for.
		case entry := <-logCh: // this is listening fpr messages on the logCh
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh: // this is listening for message on the done Ch
			break
		default: // we can actually have a default path that will be executed when nothing is recieved from both channels, then we execute the default path. this is useful when we have a non blocking select statement.
		}
	}
}

// most programming languages were designed with a single processing core in mind.	Go was born in a multi processor world.
// Waitgroups are used to synchronize goroutines together. this helps the main routine(main function) to wait for the other goroutines to finish

// Channel Basics
// when working with channels in golang we will almost work with thme in the context of goroutines. channels are designed to synchronize data transmissions between multiple goroutines.

// illustration 1.1
// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}           this is for synchronizing goroutines together

// func main() {
// 	ch := make(chan int)               this is to synchronize data flow. it is used with a built in make function, then we use a chan key word to show that we want to create a channnel and then we provide the data type thats going to flow through the channel. this channel is strongly typed. so we can only send information that are integers. same with any other type.
// 	wg.Add(2)
// 	go func() {                        this is an anonymous function. this is the recieving goroutine
// 		i := <-ch                      to recieve from the channel, the arrow points from the channel name to the recieving container(variable)
// 		fmt.Println(i)                 this goroutine recieves it and prints it out.
// 		wg.Done()
// 	}()
// 	go func() {                        this is my sending goroutine
// 		ch <- 42                       to send into the channel the arrow syntax (<-) points to the channel name. the arrow points in the direction we want the arrow to flow
// 		wg.Done()                      so this goroutine is sending the value 42 to the recieving goroutine.
// 	}()
// 	wg.Wait()
// }                                   this prints out 42

// illustration 1.2
// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func() {
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
// 	go func() {
//      i := 42
// 		ch <- i
//      i = 27                            since we are sending a copy of the variable into the channel, we can manipulate the variable and it doesnt matter.
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }                                      this also prints out 42.

// Illustration 1.3 (data that is asynchronously processed) maybe we can generate the data very quickly but it takes a lot of time to process it or maybe it takes time to generate that data but it can be processed quickly
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)                 all these routines will use this channel to communicate their data between each other.
// 	for j := 0; j < 5; j++ {             this is creating 5 set of goroutines, in each set there will be a sender and a reciever, here we are going to spun 10 goroutines
// 		wg.Add(2)
// 		go func() {
// 			i := <-ch
// 			fmt.Println(i)
// 			wg.Done()
// 		}()
// 		go func() {
// 			ch <- 42
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }                                   this prints out 42 5 times.

// illustration 1.4 (unbuffered channnels) an unbuffered channel means only one message can be in the channel at one time.
//package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	go func() {                     when the recieving goroutine is put outside the loop, the sending goroutine will send vakue 5 times but the reciever has only the capacity for one, this causes an error.
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
// 	for j := 0; j < 5; j++ {
// 		wg.Add(2)
// 		go func() {
// 			ch <- 42               the reason why its a deadlock is that this line of code pauses the execution of this goroutine until there is a space in the channel
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }                               this gives a fatal error deadlock.

// illustration 1.5 (both goroutines as senders and recievers or readers and writers.
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func() {
// 		i := <-ch                 that message is recieved in this goroutine
// 		fmt.Println(i)            and then printed out
// 		ch <- 27                  this goroutine then put another message into the channel
// 		wg.Done()
// 	}()
// 	go func() {
// 		ch <- 42                   this sends to the channel
// 		fmt.Println(<-ch)          that is recieved in this goroutine and printing the message out.
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }                               prints out 42 and 27.

// Restricting data flow
// Illustration 1.6 (most times we would want to designate particular goroutines to either reading or writing) we do that by passing a bias with a direction that we are going to work with, this done by adding variables and arguments to the goroutine.
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func(ch <-chan int) {                 this syntax shows the recieve only goroutine. arrow before chan means recieve only
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) {                 this syntax shows the send only goroutine. arrow after chan means send only
// 		ch <- 42
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }                                         this prints out 42. the above behaviour is specigic to channels

// Buffered Channels (how to solve illustration 1.4) illustration 1.7 a buffered channel is designed when the sender or the reciever is operating on another frequency to the other.
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int, 50)                    the 50 is added as a buffer to create space for the reciever to recieve values up to the buffer number.
// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		i := <-ch
// 		fmt.Println(i)
// 		i = <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) {
// 		ch <- 42
// 		ch <- 27
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()                                   the buffered channel works when the sender  or the reciever need more time to process the data, so the buffer helps create space for the information stored to stay till the data is ready to be used.
// }                                            especially when you are recieving much data in a small amount of time, ad then there is a time interval where it has a chance to process the data. the buffer helps.

// Closing Channels => when closing a channel nothing else is going to pass into it, no matter what, if you try it causes the application to panic, and there is no way to avoid the panic.
// illustration 1.8 (the right way to do the above or handle something that happens multiple times is by using a kind of looping construct)
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		for i := range ch {                      instead of processing the function and exiting immediately, we use a for range loop. we range over the channel
// 			fmt.Println(i)
// 		}
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) {
// 		ch <- 42
// 		ch <- 27
//      close(ch)                                here we close the channel showing that we are through the channel and we are not expecting more data.
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }                                             we cannot send a message into a closed channel and we cannot detect if a channel is closed before we try to send a message into it.

// For...range loops with channels
// how does for range loop work? when for a slice it executes the loop oce for evey item in the slice.
// but for a channel we need to close the channel because if we do not, we would the recieving goroutine would continue expecting

// // we can use a ,ok syntax wuth channels to pull out values(There is more than one parameter that you can pull back from the channel) illustration 1.9
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		for {
// 			if i, ok := <-ch; ok {                if ok is true then the channel is open, if ok is false then the channel is closed
// 				fmt.Println(i)                    if ok is true then we print out the message
// 			} else {                              otherwise we break
// 				break
// 			}
// 		}
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) {
// 		ch <- 42
// 		ch <- 27
// 		close(ch)
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// Select statements   illustration 2.0
//package main

// import (
// 	"fmt"
// 	"time"
// )

// const (                           this is a constant block that are declaring log level
// 	logInfo    = "INFO"
// 	logWarning = "WARNING"
// 	logError   = "ERROR"
// )

// type logEntry struct {           this is a struct that is declared that holds the time stamp for the log entry.
// 	time     time.Time              time stamp
// 	severity string                 severity of log level
// 	message  string                 the message we are trying to print out
// }

// var logCh = make(chan logEntry, 50)     here we create a log channel

// func main() {
// 	go logger()                                                                 this spins of this goroutine that is going to be the logger
// 	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
//  																			this monitors the channel for log entries that are coming in throughout the application
// 	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}              the Now function shows the current time
// 	time.Sleep(100 * time.Millisecond)                                          when we finish this sleep call here the application is turn down and the resources are reclaimed. there is no grace of ending by our self, it is torn down by itself because of the main function has finish, but most times it is better we have control over all this things. we should always have a strategy for how your goroutine should shut down. other wise there could be a resource leak which could bring our application down
// }

// func logger() {                                                              we have a loggewr here that listens to the message from the channel and prints out the below message
// 	for entry := range logCh {
// 		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
// 	}
// }
// an application is closed down as soon as the last statement of the main function completes execution.

// illustration 2.1 (ways to solve this issue) using a defer anonymous function
// package main

// import (
// 	"fmt"
// 	"time"
// )

// const (
// 	logInfo    = "INFO"
// 	logWarning = "WARNING"
// 	logError   = "ERROR"
// )

// type logEntry struct {
// 	time     time.Time
// 	severity string
// 	message  string
// }

// var logCh = make(chan logEntry, 50)

// func main() {
// 	go logger()
// 	defer func() {
// 		close(logCh)
// 	}()

// 	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

// 	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
// 	time.Sleep(100 * time.Millisecond)
// }

// func logger() {
// 	for entry := range logCh {
// 		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
// 	}
// }

// How to solve it using a select statement illustration 2.2
// package main

// import (
// 	"fmt"
// 	"time"
// )

// const (
// 	logInfo    = "INFO"
// 	logWarning = "WARNING"
// 	logError   = "ERROR"
// )

// type logEntry struct {
// 	time     time.Time
// 	severity string
// 	message  string
// }

// var logCh = make(chan logEntry, 50)
// var doneCh = make(chan struct{}) // strongly typed to a struct with no fields and this requires zero memory allocation this is called a signal only channel

// func main() {
// 	go logger()
// 	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

// 	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
// 	time.Sleep(100 * time.Millisecond)
// 	doneCh <- struct{}{} // the second curly brace is to initialize the empty struct
// }

// func logger() {
// 	for {
// 		select { //  what the select statement does it that the entire message is going to block until a message is recieved on one of the channels tha its listening for.
// 		case entry := <-logCh: // this is listening fpr messages on the logCh
// 			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
// 		case <-doneCh: // this is listening for message on the done Ch
// 			break
// 		default: // we can actually have a default path that will be executed when nothing is recieved from both channels, then we execute the default path. this is useful when we have a non blocking select statement.
// 		}
// 	}
// }
