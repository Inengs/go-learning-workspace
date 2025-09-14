package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := map[string]string{"foo": "bar", "baz": "buz"}
// 	b := a
// 	fmt.Println(a, b)
// 	a["foo"] = "qux"
// 	fmt.Println(a, b)
// }

// Creating Pointers.
// func main() {
// 	a := 42
// 	b := a                                       => what Go does is to copy whatever data is stored in a and assign it to b. so we are not actually pointing to the same memory.
// 	fmt.Println(a, b)
// }

// to prove this:
// func main() {
// 	a := 42
// 	b := a                                    => this prints: 42 42
// 	fmt.Println(a, b)                                         27 42          the second value of b didnt change with the new change of a. this shows that a and b are not pointing to the same memory.
// 	a = 27
// 	fmt.Println(a, b)
// }
// we can change this behaviour by using a pointer.
// how does the pointer notation work?
// for example:

// func main() {
// 	var a int = 42
// 	var b *int = &a                    (*int) => stands for pointer to an integer, when you need to point to a data type, you precede the figure with an asteric. (&a) the operator is called the address of operator, it helps b to point to a.
// 	fmt.Println(a, b)                 the above line says b is a pointer to an integer and i want to point it to a.
// }

// what is a pointer?
// the above prints out: 42 0xc00000a0a8 => b is holding the numerical representation for the memory address that is holding the location of a. so at that location of memory we have assigned the integer 42.
// therefore this means b isnt holding the value 42, instead it is holding the memory locagtion that is holding a's data.
// to prove that: when we ask for the address of a, by printing out &a, we get the same value as when we printed out b.
// func main() {
// 	var a int = 42
// 	var b *int = &a
// 	fmt.Println(&a, b)
// }

// Dereferencing pointers.
// just as the address of (&) operator helps us see the memory location of a variable, we can also use the dereferencing operator (*) to see what value is stored at a memory location.
// n/b: this is different from the (*int), this is declaring a pointer to data of that type, so this is a pointer to an integer. but(*b) this dereferences meaning it brings out the value in the given location.
// for example:
//
//	func main() {
//		var a int = 42
//		var b *int = &a
//		fmt.Println(a, *b)
//	}                                    => this prints out 42 42   (the same data.)
//
// but in essence what pointer do is that, in such scenanrio once there is a change in a, there will be a corresponding change in b.
// example:
//
//	func main() {
//		var a int = 42
//		var b *int = &a                                     => this prints out 42 42
//		fmt.Println(a, *b)                                                     27 27 (this means both a and b changed.) they are both pointing to the same underlying data.
//		a = 27
//		fmt.Println(a, *b)
//	}
//
// we can also change a by changing *b
// func main() {
// 	var a int = 42
// 	var b *int = &a
// 	fmt.Println(a, *b)
// 	a = 27
// 	fmt.Println(a, *b)
// 	*b = 14
// 	fmt.Println(a, *b)
// }                            => this changes the value of a to 14 to correspond with the change in b

//
// func main() {
// 	a := [3]int{1, 2, 3}                             => array of 3 values
// 	b := &a[0]                                       => pointer to the first value in the array.
// 	c := &a[1]                                       => pointer to the second value in the array.
// 	fmt.Printf("%v %p %p\n", a, b, c)                => this prints the value of a, the (memory location) pointer of b and the (memory location) pointer of c then a line break.
// }   in other languages since the memory location of c is 8 greater than b, we can substract 8 from the memory location of c and get the memory location of b, but in GO that doesn't work.
// this is called pointer arithmetic, which exists in C and C++, this results in more complex code, but because of the simplistic design of GO language, the decision was made to leave out pointer arithmetic.
// if you want to have this in your GO code, go to golang.org, then go to packages, then go to "unsafe", this will give you operations that the GO t=runtime will not check for you. If you need to know it you will learn it but generally you wont need to know how it works.

// How to create pointer types.
// we may not necessarily know what type of underlying data we are pointing to, all that is necessary is that we are able to point to such data.

// func main() {
// 	var ms *myStruct
// 	ms = &myStruct{foo: 42}
// 	fmt.Println(ms)
// }                                              => this prints out &42, which means ms is holding the address of an object that has a field with the value 42 in it.
//                                                => this is a way of initializing a variable to a pointer to an object.
// type myStruct struct {
// 	foo int
// }

// The new function: this is another way of initializing a variable to a pointer to an object. in the new function we cant use the object initialization syntax.
// func main() {
// 	var ms *myStruct
// 	ms = new(myStruct)
// 	fmt.Println(ms)
// }
//                                               => this prints out &{0}, we are only able to initialize an empty object. all of its fields will be initialized to zero.
// type myStruct struct {
// 	foo int
// }

// Working with nil
// a pointer you dont initialize, is going to be initialized for you, and it is going to hold the value <nil>
// for example:
// func main() {
// 	var ms *myStruct
// 	fmt.Println(ms)                           => this prints out <nil>
// }
// this is very important to check in your applications, because if you are accepting pointers as arguments, it is best practise to see if that pointer is a nil pointer because if it is you are going to have to handle that in a different way.

// how do we get to inner/underlying fields, for example the foo field in the above example.
// we would need to dereference the pointer.
// func main() {
// 	var ms *myStruct
// 	ms = new(myStruct)
// 	(*ms).foo = 42
// 	fmt.Println((*ms).foo)
// }

// type myStruct struct {
// 	foo int
// }                                    => this would print out 42

// the compiler can help us to make the language a bit simpler and less ugly.
// func main() {
// 	var ms *myStruct
// 	ms = new(myStruct)
// 	ms.foo = 42
// 	fmt.Println(ms.foo)
// }

// type myStruct struct {
// 	foo int
// }                                     => this would print 42 out also. this implies we want the underlying object instead of the foo field.

// How go handles variables when they are assigned one to another.
//  Types with internal pointers.
// an ARRAY of values when a particular value of a is changed, it doesnt change in the corresponding b.
// for example:
// func main() {
// 	a := [3]int{1, 2, 3}                      => this prints out: [1 2 3] [1 2 3]
// 	b := a                                                        [1 42 3] [1 2 3]
// 	fmt.Println(a, b)
// 	a[1] = 42                                                     when a was changed, b didnt change as well, because they are not pointing to the same underlying data, b points to the initial address of a
// 	fmt.Println(a, b)                                             this occurs also because the values of the array are considered intrinsic to the variable. so a holds the value of the array and the size of the array.
// }

// but this works differently for SLICES.
// func main() {
// 	a := []int{1, 2, 3}                                   => this prints out: [1 2 3] [ 1 2 3]
// 	b := a                                                                    [1 42 3] [1 42 3]
// 	fmt.Println(a, b)
// 	a[1] = 42
// 	fmt.Println(a, b)
// }                                                         but with slices, a slice is actually a projection of an underlying array. so the slice doesnt contain the data itself, the slice contains a pointer to the first element that the slice is pointing to on the underlying array.
//                                                           this means when we work with slices, the internal representation of a slice, actually has a pointer to an array. so when a is being copied to b, part of what is being copied is a pointer to a, not the underlying data itself.
// this means when you are sharing slices in your application, you are actually pointing to the same underlying data.

// another collection type that has this behaviour is a MAP.
//func main() {
// 	a := map[string]string{"foo": "bar", "baz": "buz"}         this is a map of assigning a string to string.
// 	b := a
// 	fmt.Println(a, b)                                           => this prints out: map[baz:buz foo:bar] map[baz:buz foo:bar]
// 	a["foo"] = "qux"                                                                map[baz:buz foo:qux] map[baz:buz foo:qux]
// 	fmt.Println(a, b)
// }
// This acts similar to the slices, which means it pointed to the same underlying data and such when foo was changed in a, it also changed in b.
// this just shows that passing slices and maps in your application could be detrimental, because at any point a change is done in the slice or the map, it changes it and also all of copies from that point forth.
// but this doesnt happen with primitives e.g arrays or structs.

// all assignment operations in GO are copy operations.
// slices and maps contain internal pointers
