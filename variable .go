package main

// import (
// 	"fmt"
// 	// "strconv"
// )

// func main() {
// 	var i int = 42
// 	fmt.Printf("%v, %T\n", i, i)

// 	var j string
// 	j = strconv.Itoa(i)
// 	fmt.Printf("%v, %T\n", j, j)
// }

// ways of declaring variables
// (method 1)
// var i int (used when you want to declare a variable but you dont want to initialize it yet)
// i = 42

// method 2
// var i int = 42 (Go doesnt have enough information to assign the type that you really want assigned to it) (we can easily change its type)

// method 3
// i := 42 (this will make go figure out the type) (but this is limited in the types it can guess for example: you cant make a digit float32)

// fmt.Printf("%v, %T", j, j) - this will show the value and the type of the variable.

// declaring variables at the package level.
// you cannot use (k := 99.) at package level.
// var i int
// i = 42 (you must use this) var i int = 42 (or this)

// creating a block of variables at package level.
// var (
// 	actorName string = "Elizabeth Sladen"
// 	companion string = "Sarah Jane Smith"
// 	doctorNumber int = 3
// 	season int = 11
// )
// n/b : we can have multiple variable blocks at the package level.

// How variables work when we are trying to redeclare them
// You cant declare a variable twice in the same scope.
// the variable with the innermost scope takes precedence. eg if a variable is declared at package level and also at function level; the variable at function level is printed out. (this is called shadowing)

// note : ALL VARIABLES HAVE TO BE USED.

// How to name variables.
// 1. Naming controls visibility
// lower case variable name: means this variable is scoped to this package. anything that consumes or is outside the package cant see it or work with it.
// upper case variable name: this exposes it to the outside world, when declared on the package level.
// block scope: this is only visible in the block where it is declared.
// 2. Naming conventions
// Pascal or camelCase : this means you uppercase the first letter.
// the length of the variable name should reflect the life of the variable. short variable names for short time use. longer variable names for longer use especially for package level variables.
// Capitalize acronyms: variable names with abbreviations, the abbreviations shoould be all upper case eg. theURL, theHTTP

// How to convert from one variable to another.
// var i int = 42
// 	fmt.Printf("%v, %T\n", i, i)
// 	var j float32
// 	j = float32(i) --> this is a conversion operator, float acts as a function
// 	fmt.Printf("%v, %T\n", j, j)
// but note when converting from a float32 to an int, we can loose relevant data.
// you also must explicitly convert it.

// to convert from string to other variable types we must import "strconv"
// var i int = 42
// 	fmt.Printf("%v, %T\n", i, i)

// 	var j string
// 	j = strconv.Itoa(i)
// 	fmt.Printf("%v, %T\n", j, j)

// func main() {

// }

// const OvenTime = 40

// func RemainingOvenTime(actualMinutesInOven int) int {
// 	actualMinutesInOven = OvenTime - 10
// 	fmt.Println(actualMinutesInOven)
// 	return 0
// }
