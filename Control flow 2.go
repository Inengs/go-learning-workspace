package main

// import (
// 	"fmt"
// 	"log"
// )

// func main() {
// 	fmt.Println("start")
// 	panicker()
// 	fmt.Println("end")
// }
// func panicker() {
// 	fmt.Println("about to panic")
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println("Error:", err)
// 			panic(err)
// 		}
// 	}()
// 	panic("something bad happened")
// 	fmt.Println("done panicking")

// }

// most times programs control flow from top to bottom, except maybe when we use loops to run come programs over and over again or by introducing branching logic.

// to defer one of the statements, we proceed it with the defer key word.
// func main() {
// 	fmt.Println("start")                           this prints out: start
// 	defer fmt.Println("middle")                                     end
// 	fmt.Println("end")                                              middle
// }                                               because of the defer statement
// how the defer statement works in GO, is that it actually executes any function that are passed into it after the function finishes it final statement, but before it actually returns.
// so it first calls start, moves to the next, it then sees it is deferred, so it skips it and then moves to the next and calls it and then it exits the program, and then goes back to check if there is any deferred function/statemnt and then calls it.

// if we put defer on all statements, the print out in LIFO (Last In First Out) order. this means the last one that gets deferred, will become the first called.
// func main() {
// 	defer fmt.Println("start")                  prints out: end
// 	defer fmt.Println("middle")                             middle
// 	defer fmt.Println("end")                                start
// }

// this isnt executing in context of the main function, it is executing after the main function is done, but before it is recalled.
// the most common use case of defer, is that it allows you to associate the opening of a resource and the closing of a resource next to each other.

// func main() {
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	robots, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }                                   => this is a common pattern and i will see this all the time.

// if you are making a lot of requests and opening a lot of resources in a loop, then using a defer key word may not be your best choice, because the deferred statements dont execute until the function is executed.
// so lets say you have a loop that will loop over a million resources. you are going to have a million resources open before all of them finally closes when the function closes.

// so when working with multiple resources
// you might want to explicitly close this resources when you are done with them.
// or you can delegate the processing of those resources out to another function and have that function close the resource, that way you are not keeping a bunch of resources open at one time, and wasting a lot of memory.

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	robots, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }

// func main() {
// 	a := "start"
// 	defer fmt.Println(a)                             => this prints out start, because when you defer a function like this, it actually takes the argument at the time the defer is called, not at the time the called function is executed.
// 	a = "end"
// }

// PANICKING
// in go we dont have exceptions, like a lot of applications have, because a lot of cases that are considered as exceptions are considered normal in GO.
// for example: if you try and open a file and that file doesnt exist, thats a pretty normal response, it is possible that you didnt know the file didnt exist, so in GO, we throw error values, we dont throw exceptions, because thats not considered an exception.
// but there are somethings that get a GO application to stop running, that are called exceptions in some languages in GO, we call it panic.

// Example:
// func main() {
// 	a, b := 1, 0
// 	ans := a / b
// 	fmt.Println(ans)
// }

// this prints out:
// panic: runtime error: integer divide by zero
//                                                                                  => this shows that the statement is not executable, and therefore it panics.
// goroutine 1 [running]:
// main.main()
// 	c:/Users/INENGIYE/Desktop/go-workspace/Control flow 2.go:9 +0x9
// exit status 2

// f you are going along and you are writing your own program cannot execute because of a statement that gets thrown, then it makes sense that you panic as well.
// to do this, we use the builtin panic function.
// Example:
// func main() {
// 	fmt.Println("start")
// 	panic("something bad happened")
// 	fmt.Println("end")
// }

// this prints out
// start
// panic: something bad happened                                        => it prints out the statement, tagged by panic, and then doesnt print anything afterwards.

// goroutine 1 [running]:
// main.main()
// 	c:/Users/INENGIYE/Desktop/go-workspace/Control flow 2.go:9 +0x5f
// exit status 2

// A more practical example:
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {                => this response writer gives us access to the response to this web request (what will show when we type (localhost:8080) on chrome), and then we are printing outthe string hello Go.
// 		w.Write([]byte("Hello Go!"))                                                   => the hello Go comes from this line, it is writing to the response writer above.
// 	})
// 	err := http.ListenAndServe(":8080", nil)                                           => this is an error handler. the ListenAndServe function returns an optional error object.
// 	if err != nil {                                                                    => this could occor when the port is blocked. listen and serve can fail, it wont panic, it will return an error.
// 		panic(err.Error())                                                             => we know that when this error occurs it stops the whoe operation, so what we do is we panic and print outhe error message.
// 	}                                                                                  => Go rarely will tell you whether to panic or not on an error, what it will tell us is that an error occured. it is the responsibility of you
// }

// when you type localhost:8080 on chrome, it prints out hello go, based on the code written above.

// panics dont have to be fatal. except we panic it until its go runtime and then go  runtime doesnt know what to do with a panicking application, so then it is going to kill it.
// so what do we do when our application is panicking, we can recover it.
//

// func main() {
// 	fmt.Println("start")
// 	defer fmt.Println("this was deferred")
// 	panic("something bad happened")
// 	fmt.Println("end")
// }                                              => start
//                                                   this was deferred
//                                                   panic: something bad happened

// N/B: Panics happen after deferred statements have been executed. the order of execution: main function jis executed first, then we execute any deferred statements, then we are going to handle any panics, and then we are going to handle the return value.

// why is this important.
// deferred statements that are going to close resources are going to succeed, even if your application/function panics.

// func main() {
// 	fmt.Println("start")
// 	defer func() {
// 		if err := recover(); err != nil {                  => this is called an anonymous function. the recover function on this line returns nil, if the application is not panicking. but if it isnt nil it will return the error that is causing the application to panic.
// 			log.Println("Error:", err)                     => in this logical test, we are checking to see if it is !nil, so if is not(!)nil, that means our application is panicking, and it is situationwe are simple going to log that error out.
// 		}
// 	}()                                                           => the parenthesis() on this line are making this function to execute.
// 	panic("something bad happened")
// 	fmt.Println("end")
// } this prints out:              start
//                                 2023/10/18 00:53:06 Error: something bad happened
// the end statment isn't printed out. the error was printed out through the log package.

// An anonymous function is a function that doesnt have a name. so nothing else can call this, this is defined at one point and we can call it exactly one time.
// n/b: the deferred statements doesnt take a function itself, it actually takes a function call. so you want to invoke/call the function unless things are not going to work properly.

// recover has deeper effect, if we have a bigger callstack

// func main() {                                           => this is the main function that is going to be our application entry point.
// 	fmt.Println("start")                                      then we print start (5)
// 	panicker()                                                call the panicker function (6)
// 	fmt.Println("end")                                        and then print end (7)
// }
// func panicker() { 					                   => this function called panicker.
// 	fmt.Println("about to panic")                          => this prints out about to panic.   (1)
// 	defer func() {                                         => and then it recovers from that panic (4)
// 		if err := recover(); err != nil {                  => in that recover, we are going to log out the fact that we have that error, and then we are going to log out the error message out.
// 			log.Println("Error:", err)
// 		}
// 	}()
// 	panic("something bad happened")                        => and then it is going to panic (2)
// 	fmt.Println("done panicking")                          => and then done panicking doesnt execute beacuse our application panicked (3)

// }

// this is the way it prints out: start
//								  about to panic
//								  2023/10/18 11:50:34 Error: something bad happened
// 								  end

// in the event that you are recovering from a panic, the function that actually recovers still stops execution, because it is in state where it can no longer reliably continue.
// however functions higher up the callstack, the ones that called the function that recovered, they can still continue, this is because the applicaion is in a state where the application can execute.

// this is also limiting, because you have to call the recover function, to determine what the error is, this means you are saying you can deal with it.
// so what happens if you cant deal with it, when you get the error.
// in such case, what you do is to repanic the application.

// N/B : Arguments in a deferred call are evaluated at the time the defer is executed, not at the time the called function is executed.
// when you call defer, the value of variables are going to be captured at the time that defer is executed, not at the time the function us actually executed.

// if something unplanned happens in our application, there are 2 ways our application can go, it can return an error value(normal), or it can panic.
// there are not many exceptions in a GO application that can make it to shut down, instead it prints error messages.
// for example: making an http request, and when you make the request you dont get a response from that server. it is something that happens all the time, it just means you have no response and no response is a valid response in GO.
// it just happens to be an error, because it returns a 404. in that case you return an error value you are not going to panic application.

// panics are used when you get into a situation that you cant recover from.
// dont use panic when a file cant be opened, except it is very critical.
// but you panic when there are unrecoverable events example cannot obtain TCP port for web server.

// Recover is only useful in deferred functions.
