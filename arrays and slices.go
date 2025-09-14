package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	b := a[1:]
// 	c := a[:len(a)-1]
// 	d := append(a[:2], a[3:]...)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// why do we need arrays and what are they used for.
// they help when we dont know the amount of values the application might come across when building the application.

// how do arrays look
// grades := [3]int{97, 98, 93}
//  fmt.Printf("Grades: %v", grades)

//  the [3] can be replaced by [...] this just says to the compiler that it should hold the amount of values that i am going to pass through it. but it can be any amount.

// the elements of an array are continous in memory, this makes it faster too.
// arrays are also easier to work with.

// we can declare an empty array and then fill it up later.
// var students [3]string
// fmt.Printf("Students: %v", students) => this is an empty array.

// to fill up the array, we then add values in this manner.
// students[0] = "Lisa"

// func main() {
// 	var students [3]string
// 	fmt.Printf("Students: %v\n", students)
// 	students[0] = "Lisa"
// 	students[1] = "Ahmed"
// 	students[2] = "Joy"
// 	fmt.Printf("Students: %v", students) prints out: Students: [Lisa Ahmed Joy]

// to get a particular value
// Example: to get the second value
// fmt.Printf("Student #1: %v\n", students[1])

// to review and change the value of the size of an array.
// fmt.Printf("Number of Students: %v\n", len(students))

//  an array can be made up of any type, it just has to be the same type for each array. for example, you can make an array of arrays.
// func main() {
// 	var identityMatrix [3][3]int
// 	identityMatrix[0] = [3]int{1, 0, 0}
// 	identityMatrix[0] = [3]int{0, 1, 0}
// 	identityMatrix[0] = [3]int{0, 0, 1}
// 	fmt.Println(identityMatrix)
// }                                      (an array of arrays.)

// arrays are considered values. that means when you copy arrays around it could slow your program down. especially if the array has like 1 million values.

// to help handle this we use pointers instead of assigning a value to the copy of it, we can assign the value to the pointer which points to the address of the original.
// the sign for pointer is &(it is the address of operator) it helps point to data we assign it to. so if we make a change to the pointer value it will affect the previous value of the data.

// arrays have a fixed size that has to be known at compile time.
// most common use case is to back slices.

// Slices
// a slice is initialized as a literal. syntax: []int {} empty square brackets, data type and in the curly braces, we can pass in the initialized data.
// Example: func main() {
// a := []int{1, 2, 3}
// fmt.Println(a)
//             }
// everything we can do with an array, we can do with a slice as well.
// we can check the lenght of an array/slice:
//  with this function, fmt.Printf("Length: %v\n", len(a))

// to check the capacity of a slice
// we use this function, fmt.Printf("Capacity: %v\n", cap(a))
// we can check the capacity of an array because the number of elements in the slice, doesnt necessarily match the size of the backend array.
// the slice is a projection of the underlying array.

// slices are reference types so they dont really need the pointer function: &, just declare the array (b := a) instead of (b := &a)

// if you have multiple slices pointing to the same underlying array/data, if one of those slices changes the underlying data, it could have an impact somewhere else in your application.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := []int{1, 2, 3}
// 	b := a
// 	b[1] = 5    										Prints: [1 5 3]
// 	fmt.Println(a)                                              [1 5 3]
// 	fmt.Println(b)                                              Length: 3
// 	fmt.Printf("Length: %v\n", len(a))                          Capacity: 3
// 	fmt.Printf("Capacity: %v\n", cap(a))
// }

// other ways to create slices. or called slicing operations
// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}                   prints = [1 2 3 4 5 6 7 8 9 10]
// 	b := a[:]  // slice of all elements                                  [1 2 3 4 5 6 7 8 9 10]
// 	c := a[3:]  // slice from 4th element to end                         [4 5 6 7 8 9 10]
// 	d := a[:6]   // slice first 6 elements                               [1 2 3 4 5 6]
// 	e := a[3:6]  // slice the 4th, 5th, and 6th elements                 [4 5 6]
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// }

// they are all pointing to the same underlying array/slice. so if you change anything in the array [a], it will affect b to e

// the built in make function: it takes two or three functions.
// for two functions, when making a slice.
// a := make([]int, 3) int = datatype of slice, 3 = size of slice.
// fmt.Println(a)                                  prints out: [0 0 0]
// fmt.Printf("Length: %v\n", len(a))                         Length = 3
// fmt.Printf("Capacity: %v\n", cap(a))                       Capacity = 3
// everytime we initialize a variable, we expect it to be initialized to a zero variable. this is the same with slices.

// the third value we pass into the make function is to assign the capacity.
// 	a := make([]int, 3, 100)
// 	fmt.Println(a)                                       prints out: [0 0 0]
// 	fmt.Printf("Length: %v\n", len(a))                              Length = 3
// 	fmt.Printf("Capacity: %v\n", cap(a))                            Capacity = 100
// }

// the capacity is the size of the underlying array, which can be changed.

// unlike arrays, slices dont have to have the same / fixed size all through their life, you can add more spaces and remove at any time.

// to add to a slice, we use in the built in append function. this takes two or more arguments passed into it.
// the first argument is the source slice.
// the second argument is the element you want to add to it.
// a := []int{}
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))
// a = append(a, 1)
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))

// append with multiple arguments is called a variatic function. the 2nd to last arguments are values that will be appended to the slice.
// a := []int{}
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))
// a = append(a, 1)
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))
// a = append(a, 2, 3, 4, 5, 6)
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))

// if you want to concatenate a slice to another slice. you use the spread operator which is 3 dots. this is because for you to append values to a slice they have to be of the same datatype.
// a = append(a,[]int{2, 3, 4, 5, 6}...)
// fmt.Println(a)
// fmt.Printf("Length: %v\n", len(a))
// fmt.Printf("Capacity: %v\n", cap(a))

// stack operations
// append: is to add values
// shift operation: to remove a value. to remove the first value => a[1:], to remove the last value => a[:len(a)-1].
// to remove a middle value b := append(a[:2], a[3:]...)

// as we apppend and shift the slice it changes the values of the arrays permanently.
// if you manipulate a slice in one location, it will affect all slices pointing to the underlying array. it will affect the initial array or slice from beginning.
