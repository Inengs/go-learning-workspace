package main

import (
	"fmt"
)

func main2() {
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}

// BOOLEAN TYPES.
// boolean means either true or false.
// var n bool = true
// fmt.Printf("%v, %T\n", n, n)
// uses of boolean primitives in our application.
// 1. when a user has signed for a notification when a file is generated. in order to store "true" if the want the report and the opp. for false
// 2. the more common type is as a result of logical tests
// eg. of logical tests.
// n := 1 == 1
// m := 1 == 2  (n/b : == means equal to.)
// fmt.Printf("%v, %T\n", n, n)
// fmt.Printf("%v, %T\n", m, m)  	(prints out: true, bool false, bool)
// }
// N/B: everytime you initialize a variable in GO, it has a zero value.

// NUMERIC TYPES.
// The zero value: the zero value for all numeric types is going to be zero or it is going to be the equivalent of zero for all those numeric types.
// The Integer type
// 1. Signed Integers: int: this is an integer with unspecified size. it is based on your environment. but in all environments it will never be less than 32 bits. it could be 64 or 128 bits
// 					   int8: -128 to 127
// 					   int16: -32 768 to 32 767
// 					   int32: -2 147 483 648 to 2 147 483 647
//     				   int64: -9 223 372 036 854 775 808 to 9 223 372 036 854 775 807
// 2. Unsigned Integers: there is an equivalent type of unsigned integer for every signed integer.
// 					   uint8: 0 - 255
//   				   uint16: 0 - 65 535
//    				   uint32: 0 - 4 294 967 295 ( we dont have a 64 bit unsigned integer. uint64, but we have a type byte.)
//                     byte: this is an alias for an 8 bit unsigned integer (uint8) the reason being because the uint8 is very common, that is what many datasteams use.
// 					   var n uint16 = 42
//                     fmt.Printf("%v, %T\n", n, n)
// 3. Arithmetic operations: +, -, *, /, %
// a := 10
// b := 3
// fmt.Println(a + b) prints = 13
// fmt.Println(a - b) prints = 7
// fmt.Println(a * b) prints = 30
// fmt.Println(a / b) prints = 3 (when you divide an integer by an integer you have to get an integer.) (This is called integer division.)( when you are doing division you cant change the type.)
// fmt.Println(a % b) prints = 1 (remainder).
// N/B = you cannot add two integers of different types. e.g int8 + int to make it work you will have to do a type conversion.
// var a int = 10
// var b int8 = 3
// fmt.Println(a + int(b))

// bit operators.
// a := 10 (1010 in binary)
// b := 3  (0011 in binary)
// fmt.Println (a & b) : AND operator. (prints 2 (0010); this is because it looks at the binary form of the 2 numbers and performs the AND operation (if they are both 1 == 1, if they are both 0 = 0 otherwise 0)
// fmt.Println (a | b) : OR operator. (1011 this is 11 in binary, performing the OR operation on the binary numbers)
// fmt.Println (a ^ b) : Exclusive OR operator. (XOR MEANS either one has the bit set but not both. if they are both 1 == 0) (prints 1001 which is 9)
// fmt.Println  (a &^ b) : AND NOT operator. (AND NOT MEANS only when they are both 0, we print 1. therefore this prints = 0100; which is 8)

// bit shifting: this means addinr or substacting fron the power.
//a := 8;  => 2^3
// fmt.Println(a << 3) 2^3 * 2^3 = 2^6 = 64 or 2^3+3 = 2^6 = 64
// fmt.Println(a >> 3) 2^3 / 2^3 = 2^0 = 1 or 2^3-3 = 2^0 = 1

//  floating point numbers: this are used to store decimal numbers. they follow IEEE 754 standard.
// types = float32: ± 1.18E-38 - ± 3.4E38 n/b (E means 10^)
// 		   float64: ± 2.23E-308 - ± 1.8E308
//  you cant do arithmetic operations between float64 and float32

// How to create floating point numbers
// n := 3.14 (How we define our floating point literals almost all tht time)
// n = 13.7e72 (this shows we can use exponential)
// n = 2.1e14
// fmt.Printf("%v, %T", n, n)

// Arithmetic operations available with floating point numbers.
// a := 10.2
// b := 3.7
// fmt.Println(a + b)
// fmt.Println(a - b)
// fmt.Println(a * b)
// fmt.Println(a / b)
// n/b: we have no remainder operator. no bit shifting or bit wise operators.

// Complex type: this is relatively scarce.
// Types of complex numbers: complex64 = (float32 + float32)
// `						 complex128= (float64 + float64)`
// declared in this way.
// var n complex64 = 1 + 2i
// fmt.Printf("%v, %T\n", n, n)
//
// Arithmetic operators involved with complex type
// a := 1 + 2i
// b := 2 + 5.2i
// fmt.Println(a + b)
// fmt.Println(a - b)
// fmt.Println(a * b)
// fmt.Println(a / b)

//  When you want to work with either the real or imaginary part.
//  you use the real funtion to print out real values while you use the imag function for imaginary values.
//  they work on both complex64 and complex128; for complex64 it prints results in float32. for complex128 it prints results in float64.
// eg. var n complex128 = 1 + 2i
//     fmt.Printf("%v, %T\n", real(n), imag(n))
// 	   fmt.Printf("%v, %T\n", imag(n), imag(n)) => this is the way it is written.

// to make a set of floating point numbers complex we use the complex function.
// eg. var n complex128 = complex(5, 12)
//     fmt.Printf("%v, %T\n", n, n)

// text type
// string: any UTF8 character eg. s := "This is a string"
//                                     fmt.Printf("%v, %T\n", s, s)
// we can treat it like an array which is zero based. so to bring out the third letter. we do:
//                                     s := "This is a string"
//                                     fmt.Printf("%v, %T\n", string(s[2]), s[2])

// the only operation we can do with strings is string concatenation.
// eg. s := "this is a string"
//     s2 := "this is also a string"
//     fmt.Printf("%v, %T\n", s + s2, s+ s2)

//   we can convert strings to collection of bytes, which is called a slice of bytes.
//  s := "this is a string"
//  b := []byte(s)
//  fmt.Printf("%v, %T\n", b, b)  => this prints a slice of bytes which are the UTF or ASCII values of each of the letters.
// why would we use this; a lot of functions we are going to use in GO work with byte slices. when sending it from your system to other apps or places, its good to convert it to a byte.

// Rune data type: this represents any UTF32 character. while any UTF32 character can be 32 bits long it doesnt have to be 32 bits long. so this means UTF8 characters are valid UTF32 characters.
// when we are declaring a string we use (""), when we are declaring a rune we use (''). which are single quotes
// runes are a type alias for int32
// if you are working with a data string that is encoded in UTF32 you have special functions to use. eg ReadRune
// to get more information: go to golang.org, then packages, then strings package.
//
// we dont need to memorize all data types. they normally straight forward, if we need the more complex data types. we just need to go to GO documentation.

// when you need to have control of the size of the integer or numeric type then you can care about the bit size. but normally just go for the int
