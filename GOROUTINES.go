package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	fmt.Printf("Threads: %v", runtime.GOMAXPROCS(-1))
// }

// goroutine is a lightweight execution thread in the GO programming language and a function that executes concurrently with the rest of the program.
// how to create goroutines   (illustration 1.1)
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	go sayHello()                 this addition of go immediately converts it into a goroutine, this is done by spinning of a green thread and running the sayhello function in it.
// }

// func sayHello() {
// 	fmt.Println("Hello")
// }                                this wont work because the compiler stops at only the main function.

// to fix it:                    (illustration 1.2)
//package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go sayHello()
// 	time.Sleep(100 * time.Millisecond)        this is a bad practise, but will sort it out.
// }

// func sayHello() {
// 	fmt.Println("Hello")
// }

// threads: most programming languages use OS threads, what it means is that they have an individual function callstack handed over to execute the program in that thread.
// we have a scheduler whcih help sort it out, and they can be destroyed and made easily

// typical use case of goroutine.     (illustration 1.3)
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var msg = "Hello"
// 	go func() {                                          here we use an anonymous function, it is invoked immediately and it is launched by a goroutine
// 		fmt.Println(msg)
// 	}()                                                  this works because of closure, a goroutine has access to variables in the outer scope even though the go routine is running with a completely different execution stack.
// 	time.Sleep(100 * time.Millisecond)                   this has created a slight issue because we have created a dependency between the goroutine and the main function. this is shown in illustration 1.4
// }

// illustration 1.4
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var msg = "Hello"
// 	go func() {
// 		fmt.Println(msg)
// 	}()
// 	msg = "Goodbye"
// 	time.Sleep(100 * time.Millisecond)
// }                                                         this prints out Goodbye instead of hello, this is because most times the go caller is not going to interrupt the main function until the sleep call. so even if the goroutine was inside the main func, the main func will execute first. so this can be solved as shown in illustration 1.5

// illustration 1.5
//  func main() {
// 	var msg = "Hello"
// 	go func(msg string) {
// 		fmt.Println(msg)
// 	}(msg)                                  this helps detach the go routine from the redeclaration of msg on the next line. this is generally the way we would want to pass data into our goroutine
// 	msg = "Goodbye"
// 	time.Sleep(100 * time.Millisecond)      but this isnt fully best practise because we are using the sleep call and what this means is that we are binding the program to the real world clock cycle thereby slowing down its clock time. some alternatives to sleep call are written in the next illustration.
// }                                        this is done by using arguments to declare the function for the goroutine

// illustration 1.6 (waitgroup) a waitgroup is designed to synchronize multiple waitgroups together.
// example: synchronizing the main function go routine to the goroutine in the main function

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}	here we create a variable wg and pull out the sync package and then create an object WaitGroup and then the curly braces{} is to initialize it

// func main() {                we have two goroutines, the one executing the main function and then the one on line 113 and we want to synchronize both of them
// 	var msg = "Hello"
// 	wg.Add(1)                   we do that by telling the waitgroup that we have another goroutine we want it to synchronize to, and then since it initializes from 0 we then add 1 since it is one waitgroup we are syncronizing with another.
// 	go func(msg string) {
// 		fmt.Println(msg)
// 		wg.Done()               this is written for the goroutine to tell the waitgroup that it is done
// 	}(msg)
// 	msg = "Goodbye"
// 	wg.Wait()                   this is written to exit or leave the application, this is done by using the wait method.
// }

// the next example we have multiple goroutines working on data. illustration 1.7
// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}                 here i am creating a waitgroup
// var counter = 0                           here we initialize a countergroup at 0,

// func main() {
// 	for i := 0; i < 10; i++ {                here we have about 20 goroutines
// 		wg.Add(2)                            everytime we run the for loop we add 2 goroutines
// 		go sayHello()                        this is one of them
// 		go increment()                       this is the other.
// 	}
// 	wg.Wait()                                here we have a wait method call.
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v/n", counter)       here we print out hello and the counter value
// 	wg.Done()                                this shows we are done, this is a done method
// }

// func increment() {
// 	counter++                                here we just increment the counter by one.
// 	wg.Done()
// }                             this prints out this: it is not synchronized.   Hello #2/nHello #1/nHello #4/nHello #5/nHello #7/nHello #6/nHello #8/nHello #9/nHello #10/nHello #0/n

// to fix the above and synchronize it we use a mutex. as shown in illustration 1.8, A mutex is a lock that the application is going to honour
// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var counter = 0
// var m = sync.RWMutex{}         this is an RWmutex, a read-write mute. any amount of things can read it, but only one thing can write to it at a time and if any thing is reading then we cant write to it at all. we can have an infinite number of readers but only one reader.

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	m.RLock()                           here we do a readlock on the printf function
// 	fmt.Printf("Hello #%v/n", counter)
// 	m.RUnlock()                         and here we release it by a RUnlock(R Unlock)
// 	wg.Done()
// }

// func increment() {
// 	m.Lock()                             here we call the lock method
// 	counter++                            here we increment the counter variable
// 	m.Unlock()                           and here we unlock it.
// 	wg.Done()                                     this prints out Hello #0/nHello #1/nHello #1/nHello #1/nHello #1/nHello #1/nHello #1/nHello #1/nHello #1/nHello #1/n : - this is because the sayhello runs completely before the counter incement in the increment function has chance to run. this is solved in the next illustration
// }                                              a simple mutex is either locked or unlocked, so if a mutex is locked and something want to manipulate that value the thing has to wait for it to be unlocked, or it can obtain the mutex lock itself, we can protect parts of our code and only one part can be manipulated at a time.

// Illustration 1.9
// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var counter = 0
// var m = sync.RWMutex{}

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		m.RLock()                    here we put the readlock in the main function before each goroutine executes
// 		go sayHello()
// 		m.Lock()                     this is the same
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v/n", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	m.Unlock()
// 	wg.Done()                                     this has caused a new problem because this has destroyed concurrency and parralelism. because this has forced everything to run in a single thread. as a matter of fact this application will perform better without than with goroutines as i have done.
// }

// lets talk about GOMAXPROCS Package : this shows the number of threads available in the application. what the system does is to give you the number of threads equal to the number of cores available in the machine. no
//  func main() {
// 	runtime.GOMAXPROCS(1)                                      with this we can cange the amount of threads used.nor we have a truly cncurrent application with no parralelism.
// 	fmt.Printf("Threads: %v", runtime.GOMAXPROCS(-1))          the -1 just helps us interrogate the number of threads available
// }
// one advise is on operating system thread per core, but our application might get faster by increasing GOMAXPROCS beyond that value. if it goes too high, the performance can peak and start to fall back off. so when producing an app, it will be good to test with different values of GOMAXPROCS to see where the app will run best.
// when running locally in an app, we can use multiple cores
// tools we have that are available for parallel and concurrent programming
// goroutines help us create highly efficient and concurrent applications

//Creating goroutines
// Synchronization: WaitGroups and Mutexes

// goroutines will allow our applications work on multiple things at the same time

// Parallelism: concurrency is the ability of an application to work on multiple things at the same time, it doesnt mean it can work on them at the same time, it means it has multiple things it can be doing. parallelism is how we take the Go applications and enable them to work on those concurrent calculations in parallel.

// Best Practices
// 1. when working on a library, be careful of creating GOROUTINES, it better the consumer controls concurrency of the library, let the person decidse when to use
// 2. when creATING A GOROUTINE know how it is going to end. avoids subtle memory leaks. when it ages too much it can cause your app to crash.
// 3. Check for race conditions at compile time. how to fix issues without fully compiling it, you can add -race flag to your go run command. it is good wh
