package main

// import (
// 	"fmt"
// )

// func main() {
// 	var i interface{} = 1
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is an int")
// 		break
// 		fmt.Println("This will print too")
// 	case float64:
// 		fmt.Println("i is a float64")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("i is another type")
// 	}
// }

// If statements: we are going to spend most of the time trying to generate boolean results.
// Example: this is one of the simplest forms.
// func main() {
// 	if true {
// 		fmt.Println("The test is true")
// 	}
// }

// N/B you are never allowed a single line block evaluate as an if statement, it must be kept inside curly braces even if it is one line of code, you must use curly braces.
//
// another style of if statement: the initializer syntax.

// func main() {
// 	statePopulations := map[string]int{
// 		"California":   39250017,
// 		"Texas":        27862596,
// 		"Florida":      20612439,
// 		"New York":     19745289,
// 		"Pennsylvania": 12802503,
// 		"Illinois":     12801539,
// 		"Ohio":         11614373,
// 	}
// 	if pop, ok := statePopulations["Florida"]; ok {
// 		fmt.Println(pop)
// 	}
// }

//pop, ok := statePopulations["Florida"]; ok { this is the initializer.
// this creates a map called ok. pop is the variable.
// the second ok is the condition that will be evaluated.
// this creates a boolean result by interrogating the map.
// the pop variable is scoped to the block of the if statement.

// Comparison operators: >, <, == : Greater than, Less than or Equal to.
// func main() {
// 	number := 50
// 	guess := 30
// 	if guess < number {
// 		fmt.Println("Too low")
// 	}
// 	if guess > number {
// 		fmt.Println("Too high")
// 	}
// 	if guess == number {
// 		fmt.Println("You got it ")
// 	}
// }

// <= : less than equal to operator, >= : greater than equal to operator, != : not equal to
// fmt.Println(number >= guess, number <= guess, number != guess)
// These are the Six different Comparison operators.

// When working with string types, what we will most likely work with is equality(==) and non-equality(!=)

// to validate a user's input to be sure they dont put in a number like (-5), we use logical operators(AND, OR AND NOT)
// example: this conducts two tests, (||) stands for "or" operation. one of them needs to be true.
//  if guess < 1 || guess > 100 {
// 	fmt.Println("The guess must be between 1 and 100!")
// }
// example 2: this conducts two tests, (&&) stands for "and" operation. this says both cases have to be true.
// if guess >= 1 && guess <= 100 {
// how to enable the user to enter the guess (Similar to input function in Python.)

// Not operator: this takes a boolean and flips it to the other side. (!) this stands for "not" operation.
// fmt.Println(!true) prints out: false

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	number := 50
// 	guess := -5
// 	if guess < 1 || guess > 100 {
// 		fmt.Println("The guess must be between 1 and 100!")
// 	}
// 	if guess >= 1 && guess <= 100 {
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess > number {
// 			fmt.Println("Too high")
// 		}
// 		if guess == number {
// 			fmt.Println("You got it!")
// 		}
// 		fmt.Println(number >= guess, number <= guess, number != guess)
// 	}
// 	fmt.Println(!true)
// }

// Short circuiting: this occurs when one part of an "OR" test returns true, its not going to evaluate the rest of the "OR" test and therefore it executes the instruction for the test instead of checking the remaining other conditions.
// this also occurs for an "AND" test, because if the first condition returns false, the compiler will not bother wuth the remaining other conditions.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	number := 50
// 	guess := 30
// 	if guess < 1 || returnTrue() || guess > 100 {
// 		fmt.Println("The guess must be between 1 and 100!")
// 	}
// 	if guess >= 1 && guess <= 100 {
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess > number {
// 			fmt.Println("Too high")
// 		}
// 		if guess == number {
// 			fmt.Println("You got it!")
// 		}
// 		fmt.Println(number >= guess, number <= guess, number != guess)
// 	}
// 	fmt.Println(!true)
// }

// func returnTrue() bool {              => this is a function called returnTrue that will return a boolean result. To know that it executed, we then print returning true.
// 	fmt.Println("returning true")        => so anytime the return function is called, it prints "returning true"
// 	return true
// }

// short circuiting occurs in the above code if guess is equal to (-5), since the guess is less than 1, it automatically forgets the return true condition and also the guess > 100 condition and just goes on to print the statement.

// the if else statements: this is done when we need either of the statements to be executed.
// func main() {
// 	number := 50
// 	guess := 30
// 	if guess < 1 || guess > 100 {
// 		fmt.Println("The guess must be between 1 and 100!")
// 	} else {
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess > number {
// 			fmt.Println("Too high")
// 		}
// 		if guess == number {
// 			fmt.Println("You got it!")
// 		}
// 		fmt.Println(number >= guess, number <= guess, number != guess)
// 	}
// }

// in a situation where we want multiple tests to be carried out: we need to add the clause called "else if"
// func main() {
// 	number := 50
// 	guess := 200
// 	if guess < 1 {
// 		fmt.Println("The guess must be between 1 and 100!")
// 	} else if guess > 100 {
// 		fmt.Println("The guess must be less than 100!")
// 	} else {
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess > number {
// 			fmt.Println("Too high")
// 		}
// 		if guess == number {
// 			fmt.Println("You got it!")
// 		}
// 		fmt.Println(number >= guess, number <= guess, number != guess)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	number := 50
// 	guess := 200
// 	if guess < 1 {
// 		fmt.Println("The guess must be between 1 and 100!")
// 	} else if guess > 100 {
// 		fmt.Println("The guess must be less than 100!")
// 	} else {
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess > number {
// 			fmt.Println("Too high")
// 		}
// 		if guess == number {
// 			fmt.Println("You got it!")
// 		}
// 		fmt.Println(number >= guess, number <= guess, number != guess)
// 	}
// }

// floating point numbers are approximations of decimal values, they are not exact representation of decimal values, so when doing comparison operations between decimal numbers, if else may not be a good idea.
// equality tests is not generally a good idea with GO.
// Example: this will print out "these are different" instead of bringing out "they are the same"
// func main() {
// 	myNum := 0.123
// 	if myNum == math.Pow(math.Sqrt(myNum), 2) {       => this is an operation which takes the square of myNum, and then takes the square root. this is supposed to give the exact number.
// 		fmt.Println("They are the same")
// 	} else {
// 		fmt.Println("These are different")
// 	}
// }

// to solve this; we generate an error value and check to see if the error value is less than a certain threshold.
// func main() {
// 	myNum := 0.123
// 	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {      => this means that myNum will divide the value gotten from above and get value2,
// 		fmt.Println("They are the same")                                 then 1 will be substracted from value2 to get value3
// 	} else {                                                             if the difference (value3) is less than 0.001, then we print "They are the same" else we print "they are different"
// 		fmt.Println("These are different")
// 	}
// }

// SWITCH STATEMENTS: this is a special purpose if statement.
//
// func main() {
// 	switch 2 {                                     => no. 2 is called tag.
// 	case 1:                                        => then 1 is compared to the tag, if true the statement is printed.
// 		fmt.Println("one")
// 	case 2:                                        => the same thing happens here
// 		fmt.Println("two")
// 	default:                                       => the default case then occurs, when none of the above cases pass.
// 		fmt.Println("not one or two")
// 	}
// }       => this prints out "2"

// note: once one of the cases passes, the process stops, just as it occurs with if statements.
// its an easy way of compare one variable to multiple possible values of that variable.

// to compare a range : we can compare a range of values per case.
// for example:
// func main() {
// 	switch 2 {
// 	case 1, 5, 10:
// 		fmt.Println("one, five or ten")
// 	case 2, 4, 6:
// 		fmt.Println("two, four or six")
// 	default:
// 		fmt.Println("another number")
// 	}
// }

// note: you cant have the same test in different cases.

// we can also use an initializer, instead of just a tag.
// func main() {
// 	switch i := 2 + 3; i {
// 	case 1, 5, 10:
// 		fmt.Println("one, five or ten")
// 	case 2, 4, 6:
// 		fmt.Println("two, four or six")
// 	default:
// 		fmt.Println("another number")
// 	}
// }

// we can also use a tagless syntax: this is arguably more powerful than the tagged syntax although it is slightly more verbose.
// for example.
// func main() {
// 	i := 10
// 	switch {
// 	case i <= 10:                                          => we can use both logical and camparison operators.
// 		fmt.Println("less than or equal to ten")
// 	case i <= 20:
// 		fmt.Println("less than or equal to twenty")
// 	default:
// 		fmt.Println("greater than twenty")
// 	}
// }
// when using tagless syntax (unlike the tagged syntax) they can overlap.
// they are all part of a block so each of them dont need to be put into curly braces.
// case, default key words and the curly braces delimate the end of a block.
// we have an implicit break in between our cases, instead of an implicit fall through.

// what then happens if we want our case to fall through
// we can actually use the keyword "fallthrough"
// func main() {
// 	i := 10
// 	switch {
// 	case i <= 10:
// 		fmt.Println("less than or equal to ten")
// 		fallthrough
// 	case i <= 20:
// 		fmt.Println("less than or equal to twenty")
// 	default:
// 		fmt.Println("greater than twenty")
// 	}
// }         => This prints out less than or equal to ten and also less than or equal to twenty.

// N/B: fallthrough is logicless. so this means when we say fallthrough the next case, it will do that and execute the next case, whether it is true or false.
// func main() {
// 	i := 10
// 	switch {
// 	case i <= 10:
// 		fmt.Println("less than or equal to ten")
// 		fallthrough
// 	case i <= 20:                                                       => even if this is >= 20, it will still print the statement, when 10 is not >= 20.
// 		fmt.Println("less than or equal to twenty")
// 	default:
// 		fmt.Println("greater than twenty")
// 	}
// }

// TYPE SWITCH
// example: the tag is typed to type "interface", we can assign anything to an interface, both data types and collection types.
// func main() {
// 	var i interface{} = 1
// 	switch i.(type) {                                    => we use the .(type) to tell go to pull the actual underlying type of that interface, and use that data for whatever we are doing next. we use it for type switching and other things.
// 	case int:
// 		fmt.Println("i is an int")
// 	case float64:
// 		fmt.Println("i is a float64")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("i is another type")
// 	}
// }                           => this prints out i is an int because 1 is an int.

// arrays are not evaluated as the same, they have to have the same datatype and the same size for them to be considred the same array. eg. [3]int{} is only of array type [3]int.
// func main() {
// 	var i interface{} = [3]int{}
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is an int")
// 	case float64:
// 		fmt.Println("i is a float64")
// 	case string:
// 		fmt.Println("i is a string")
// 	case [3]int:
// 		fmt.Println("i is a [3]int{}")
// 	default:
// 		fmt.Println("i is another type")
// 	}
// }

// if you dont want to execute all the statements under the case, you can use the break keyword.
// example:
// func main() {
// 	var i interface{} = 1
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is an int")
// 		break
// 		fmt.Println("This will print too")
// 	case float64:
// 		fmt.Println("i is a float64")
// 	case string:
// 		fmt.Println("i is a string")
// 	default:
// 		fmt.Println("i is another type")
// 	}
// }
