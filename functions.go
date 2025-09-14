package main

// import (
// 	"fmt"
// )

// func main() {
// 	g := greeter{
// 		greeting: "Hello",
// 		name:     "Go",
// 	}
// 	g.greet()
// 	fmt.Println("The new name is:", g.name)
// }

// type greeter struct {
// 	greeting string
// 	name     string
// }

// func (g *greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = " "
// }

// Basic syntax
// the way GO is structured is that, you always have to have an entry point in the application. the entry point of GO application is always in package main.
// and within that package main, you have to have a function called main that takes no parameters and returns no value.

// for example:
// package main         => main package

// import (
// 	"fmt"
// )
//
// func main() {                                     => main function, no parameters, no return value.
// 	fmt.Printf("%v, %T\n", "I am a boy", 12)         => the application actually starts on this line.
// }

// things to know about functions.
// they start with a func keyword. and that describes a function we are going to create.

// after that we have the name of the function we are going to create. in the above example the name is "main". this name follows the naming convention of any other variable in GO. pascal case or camel case for the name of your function.
// upper case or lower case first letters determines the visibility. upper case first letter, is going to be public from the package so anything else can use it, and a lower case first letter is going to be kept internal.

// the next thing is the mashed parenthesis, which is compulsory syntax, although there are no parameters.

// the actual content of the function is contained in curly braces. the opening curly brace has to be on the same line as the func keyword. and then the closing curly brace generally has to be on hiw own. (there are some exceptions)
// the execution path is static,  we cant influence what a function is doing from the outside.

// Parameters: this is to pass information into it.
// func main() {
// 	sayMessage("Hello GO!")                        here main function calls into another function called sayMessage. when you want to call that parameter, we have to pass in the value for that parameter and thats called an ARGUMENT.
// }                                               so here we pass in the argument Hello GO!
// func sayMessage(msg string) {                   this sayMessage takes in the parameter msg thats of type string. whenever you are defining a function that takes parameters, the parameters go between this 2 parenthesis and they are described like any other variable declaration, except you dont need the var keyword.
// 	fmt.Println(msg)                               and inside our function we are printing out the message that gets passed in.
// }                                               this msg parameter isnt special in any other way than the fact that it is passed into the function. it is treated as a local variable, just as any other variable we might create.

// we can pass multiple variables in:
// func main() {
// 	for i := 0; i < 5; i++ {
// 		sayMessage("Hello Go!", i)                            when we call the function, we are going to enter a value for each one of those parameters, they have to be in the same other that they are declared. the first one passed is a string for the msg variable, then we pass in the i that is going to be a loop counter which is declared above and it is going to be put into the idx variable.
// 	}
// }
// func sayMessage(msg string, idx int) {                     here we can see the sayMessage fucntion, holds unto 2 parameters. the second parameter is an index (idx) and tis is of type integer here.
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is", idx)
// }
// we can pass in any amount of parameters, we are just going to provide a comma to seperate them.
// the above prints out:
// Hello Go!
// The value of the index is 0
// Hello Go!
// The value of the index is 1                  the value of index changes everytime the function is called because of the loop counter.
// Hello Go!
// The value of the index is 2
// Hello Go!
// The value of the index is 3
// Hello Go!
// The value of the index is 4

// when you are passing in a function you are going to pass in multiple parameters of the same type. we may have syntax like this:
// func main() {
// 	sayGreeting("Hello", "Stacey")                                          instead of the normal sayHello, we will have a sayGreeting function
// }
// func sayGreeting(greeting string, name string) {                         this sayGreeting, possesses a greeting string and a name string.
// 	fmt.Println(greeting, name)                                             this prints Hello Stacey
// }
// since the types are the same, we dont need to make it this verbose.
// it can be written like this:
// func main() {
// 	sayGreeting("Hello", "Stacey")
// }
// func sayGreeting(greeting, name string) {               we can write it like this. it gets the same syntax.
// 	fmt.Println(greeting, name)
// }

// the difference between passing in values and passing in pointers.
// func main() {
// 	greeting := "Hello"
// 	name := "Stacey"                                         => this prints out the same thing Hello Stacey.
// 	sayGreeting(greeting, name)
// }
// func sayGreeting(greeting, name string) {
// 	fmt.Println(greeting, name)
// }

// func main() {
// 	greeting := "Hello"
// 	name := "Stacey"                                                              => the name variable is passed by value.
// 	sayGreeting(greeting, name)
// 	fmt.Println(name)
// }                                                                              n/b: the GO runtime will copy the name variable and provide it to the sayGreeting function.
// func sayGreeting(greeting, name string) {
// 	fmt.Println(greeting, name)
// 	name = "Ted"                                                                   => the name variable is changed in the function.
// 	fmt.Println(name)                                                              => then it is printed out on this line.
// }
// this prints out : Hello Stacey
//                   Ted
//                   Stacey
// this is a safe way of passing data into a function, we can be assured by passing by value , that the data is not going to be changed. but when passing pointers, it is different.
// when passing pointers to our variables around in our application, this means instead	of working with a copy of name variable, we are working with the pointer to the name variable.
// this is how it looks.
// func main() {
// 	greeting := "Hello"
// 	name := "Stacey"
// 	sayGreeting(&greeting, &name)
// 	fmt.Println(name)
//                                                                               this prints out: Hello Stacey
// }                                                                                              Ted
// func sayGreeting(greeting, name *string) {                                                     Ted
// 	fmt.Println(*greeting, *name)
// 	*name = "Ted"
// 	fmt.Println(*name)
// }
// here we see not only have we changed the variable in the scope of the function, but in the calling scope as well. so by passing a pointer we have infact manipulated the parameter that we passed in.
// why would we want to do this?
// 1. a lot of times our functions do need to act on the functions passed into them, and so passing in pointers is really the only way to do that.
// 2. a lot of times passing in a pointer is a lot more efficient than passing in a whole value. pessing in copies and passing in pointers will be the same in terms of performance, however if you are passing in a large data structure, then passing in the value of that data structure is going to cause that entire data structure to be copied everytime. so in that case you might decide to pass in a pointer simply for performance benefits.
// we need to be careful when passing pointers because you can't inadvertently change the value, and so this might cause issues.

// if you are working with maps or slices, then you dont really have these options because they have internal pointers to there underlying data, then they are always going to act like you are passing pointers in.

// variatic parameters:
// func main() {
// 	sum(1, 2, 3, 4, 5)                                                  => and then we are passing in the numbers 1 - 5
// }
// func sum(values ...int) {                                            => here we create a generic sum function. in this line we have one variable called values and then its type is preceded with 3 dots(...), the 3 dots has told the go runtime to take all of the last arguments that are passed in and wrap them up into a slice, that has the name of the variable (values).
// 	fmt.Println(values)                                                    inside the sum function, we print out what the values object is, so we can see it.
// 	result := 0
// 	for _, v := range values {                                             since it acts like a slice, we can use the for loop and range over those values
// 		result += v                                                        for every value, this expression is carried out.
// 	}
// 	fmt.Println("The sum is ", result)                                     then we print out the result.
// }             this prints out: [1 2 3 4 5]
//                                The sum is 15
// when using a variatic parameter, you can only have one and it has to be the last one. for example:
// func main() {
// 	sum("the sum is", 1, 2, 3, 4, 5)
// }
// func sum(msg string, values ...int) {                               => this still gives the same result.
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	fmt.Println(msg, result)
// }                                but the runtime will not allow us to put the message parameter after the values parameter(last). this is because the run time doesnt have the ability to understand where the variatic parameters end.
// this means when using variatic parameters you can only have one and they have to be at the end.

// Return values: this is when we use our functions to do some work and then return the value back to the calling function or main function.
// func main() {
// 	s := sum(1, 2, 3, 4, 5)                                     we can catch the return value by declaring a variable and asssigning it to the result of the function. so s is going to be an integer type.
// 	fmt.Println("The sum is ", s)
// }
// func sum(values ...int) int {                                the second int on this line is to assign the value of the return value.
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return result                                               here we use the return keyword, and then return the result of the variable that we have building up in the course of the function.
// }

// another ability that GO has that is pretty rare in other languages is the ability to return a local variable as a pointer.
// func main() {
// 	s := sum(1, 2, 3, 4, 5)                               s is now a pointer
// 	fmt.Println("The sum is ", *s)                        the s is now dereferenced.
// }
// func sum(values ...int) *int {                         we return a pointer to an integer (*int)
// 	fmt.Println(values)
// 	result := 0                                           the result declared here is declared on the execution stack of the sum function.
// 	for _, v := range values {
// 		result += v
// 	}
// 	return &result                                        here we return the address of the result.
// }
// in GO language when it recognizes that you are returning a value thats generated on the local stack, it automatically promotes this value for you to be on the shared memory on this computer, what is also called the heap memory.

// using named return values:
// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is ", s)
// }
// func sum(values ...int) (result int) {             we have a set of parenthesis now and then we have a name for the return value and a type for it (result int), this is basically syntactic sugar for declaring a result variable.
// 	fmt.Println(values)
// 	for _, v := range values {
// 		result += v                                   this variable is available in the scope of our sum function.
// 	}                                                 the value is going to be implicitly returned.
// 	return                                            we dont have have to specify the name of the returned variable.
// }
// this is not realy done in GO, probably because it is not really easy to read. when having longer functions, named functions can actually be more confusing than less confusing, this is because to see the value that is going to be returned we have to go to the top of the function.

// Doing multiple return values from a function.
//func main() {
// 	d := divide(5.0, 3.0)
// 	fmt.Println(d)
// }
// func divide(a, b float64) float64 {                    => this is a simple divide function that takes in two parameters, that are float64. its going to divide them and it is going to return a float64 value.
// 	return a / b
// }                                                         this prints out: 1.66666666666666667

// what if we divide 5 by 0
// func main() {
// 	d := divide(5.0, 0.0)
// 	fmt.Println(d)
// }
// func divide(a, b float64) float64 {
// 	return a / b
// }                                                            => this prints out a +Inf, but to be able to handle invalid input values. some languages could panic the app or create exceptions. so to do that; we create an if statement.

// func divide(a, b float64) float64 {
// 	if b == 0.0 {
// 		panic("Cannot provide zero as second value")
// 	}                                                           =>this application cannot continue because the panic stops the application, this means this application wont continue once a person puts a value of b = 0.
// 	return a / b                                                  it is right to assume that we might pass 0 in the application occassionally, what we wanna do is return an error back, letting the calling function know that something they asked it to do, hasnt been done properly.
// }

// func main() {
// 	d, err := divide(5.0, 0.0)                                  when we are recieving multiple values out of a function call, we actually have a common delemated list of those return values.
// 	if err != nil {                                             here we have the standard tests to see if error is not equal to nil. if it isnt equal to nil, then we need to process the error.
// 		fmt.Println(err)
// 		return                                                  here we return from our main function, so we exit our apllication.
// 	}
// 	fmt.Println(d)                                              here we print out the result of our application
// }
// func divide(a, b float64) (float64, error) {                 inside the function that can generate an error, you are going to return the expected value and then error as the second parameter.
// 	if b == 0.0 {                                               this checks for the error conditions, you will return from the function as soon as possible with the error value if an error is present.
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil                                           here we return the result of our calculation and then a nil error.
// }                  => this prints out cannot divide by zero.
// func main() {
// 	d, err := divide(5.0, 3.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }
// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// Anonymous function:
// func main() {
// 	func() {                                            => this is known as an anonymous function. starts with a func keyword, then the parenthesis for the parameters, then the curly braces, but it doesnt have a function name, this is what makes it anonymous.
// 		fmt.Println("Hello Go!")                        =. this is the basic format of functions when you are not working with them in the traditional scope, but instead you are working with them as types.
// 	}()                                                 => the parenthesis on this line, invokes this function, so this is an immediately invoked function. this makes us to both define and execute the function at the same time.
// }
// you can create variables in the anonymous function, that will only exist in that scope, and not in the package scope or other external scopes.

// when creating for loops in an anonymous function:
// func main() {
// 	for i := 0; i < 5; i++ {                => this gives a bit of an odd behaviour when working with asynchronous behaviour
// 		func() {                            if this is executed asynchronously we could have an issue with the counter, it may not stop and we might have odd behaviour, so the better way to do it is shown below:
// 			fmt.Println(i)                  we have access to this i variable because we are in the scope of the main function, so inner functions can take advantage of variables that are in the outer scope.
// 		}()
// 	}
// }

// this is better practise:
// func main() {
// 	for i := 0; i < 5; i++ {                       => what this does is we are not reading from th outer scope anymore, we are passing that into the function execution and that way even if this is running asynchronously, we are going to print out the value correctly. this way changes in the outer scope are not reflected in the inner scope.
// 		func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}
// }

// Functions as types: they can be passed around as variables, they can be passed as arguments in the functions, you can get them as return values, really anything we can do with any other type we can do with functions
// functions as variables:
// func main() {
// 	f := func() {                                        => we declared an anonymous function and assigned it to a variable f.
// 		fmt.Println("Hello GO!")
// 	}
// 	f()                                                  => this is then free to pass around the application because of it is a variable.
// }

// what do we use as the signature: this is how we find out.
//func main() {
// 	var f func() = func() {
// 		fmt.Println("Hello GO!")
// 	}
// 	f()
// }

// func main() {
//     var divide func(float64, float64) (float64, error)                     => we declare a function signature for a divide function, that takes in 2 floats and returns a float and an error.
//     divide = func(a,b float64) (float64, error) {                          => we then initialize the variable to an anonymous function that takes a and b
// 		if b == 0.0 {
// 			return 0.0, fmt.Errorf("Cannot divide by zero")
// 		} else {
// 			return a / b, nil
// 		}
// 	}
// 	d, err := divide (5.0, 3.0)                                                => the difference between this and declaring this globally is that, if declared at package level, it will give a compiler error.
// 	if err != nil {                                                            => when working with functions as variables make sure they are defined before you execute them.
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// Methods
// func main() {
// 	g := greeter{                                                here we declare a greeter struct
// 		greeting: "Hello",
// 		name:     "Go",
// 	}
// 	g.greet()                                                    and then here we call the function and then proceed it with the struct we have above. this is how we do method invocation we call the method just like when we are accessing a field, except we have parenthesis where we can pass in arguments.
// }

// type greeter struct {                                         this is a struct called greeter, this has 2 fields; greeting and name.
// 	greeting string
// 	name     string
// }

// func (g greeter) greet() {                                     (g greeter) => This is what causes this to be a method, a method is a function that is executing in a known context. and a know context in GO is any type.
// 	fmt.Println(g.greeting, g.name)                               when we call the greet method, the greet method is going to get a copy of the greeter object and thats going to be given the name g in the context of this method.
// }                                                              when we print out we can get the fields on the collection type in this case the struct, here we print the greeting and the name.
// it provides a context of where the function is executing in.
// when we use this syntax, it specifies greeter as a value type, we dont have a pointer here. This is what is called a value reciever. the recieved object in this greet method is the value "greeter" this means when working with the struct we are working with a copy and not the main struct or a pointer.
// we are operating on the copy of the greeter object and not the greeter object itself. so everytime we are working with a large struct, everytime we invoke a method open a new copy of the struct.
// but if we make it a pointer, we can manipulate the underlying data.

// func main() {
// 	g := greeter{
// 		greeting: "Hello",
// 		name:     "Go",
// 	}
// 	g.greet()
// 	fmt.Println("The new name is:", g.name)                        => this prints out Hello Go
// }                                                                                  The new name is: Go

// type greeter struct {
// 	greeting string
// 	name     string
// }

// func (g greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = " "
// }

// func main() {
// 	g := greeter{
// 		greeting: "Hello",
// 		name:     "Go",
// 	}
// 	g.greet()                                                          This prints out: Hello Go
// 	fmt.Println("The new name is:", g.name)                                             The new name is:
// }                                                                   here the underlying data was manipulated.

// type greeter struct {
// 	greeting string
// 	name     string
// }

// func (g *greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = " "
// }
