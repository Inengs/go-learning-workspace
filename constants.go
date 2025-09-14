package main

// import (
// 	"fmt"
// )

// const (
// 	isAdmin = 1 << iota
// 	isHeadquarters
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmerica
// )

// func main() {
// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
// 	fmt.Printf("%b\n", roles)
// 	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
// 	fmt.Printf("Is HR? %v", isHeadquarters&roles == isHeadquarters)
// }

// Naming conventions
// we must use the const function to start it
// if the first letter is capital letter, it can be exported.
// if the first letter is small letter, it cant be exported. using camel case.

// Typed constant
// const myConst int = 42
// fmt.Printf("%v, %T\n", myConst, myConst) => written in this manner

//  we are not allowed to change the value of a constant.
//  you cannot set your constant to something that has to be determined at runtime.

// Constants can be made of any of the primitive types that we talked about in the last video.
// const a int = 14
// const b string = "foo"
// const c float32 = 3.14
// const d bool = true
// fmt.Printf("%v\n", a)
// fmt.Printf("%v\n", b)
// fmt.Printf("%v\n", c)
// fmt.Printf("%v\n", d)

// collection types are inherently mutable.

// Constants can be shadowed.
// import (
// 	"fmt"
// )

// const a int16 = 27

// func main() {
// 	const a int = 16
// 	fmt.Printf("%v, %T\n", a, a) this prints => 16, int

// Comstants can be added to variables.
// const a int = 42
// var b int = 27
// fmt.Printf("%v, %T\n", a+b, a+b) this prints 69, int

// the compiler can infer the datatype for us.
// const a = 42
// fmt.Printf("%v, %T\n", a, a) this prints 42, int.

// but you can add the constant with an inferred type to any datatyped variable
// eg. const a = 42
// 	   var b int16 = 27
//     fmt.Printf("%v, %T\n", a+b, a+b) => this prints 69, int16

// innumerated constants: this is most commonly applied at package level.
// import (
// 	"fmt"
// )
// const a = iota
// func main() {
// 	fmt.Printf("%v, %T\n", a, a) => this prints 0, int
// iota is a counter that can be used when we are creating innumerated constants.

// constants at package level can be made into constant blocks
// import (
// 	"fmt"
// )

// const (
// 	a = iota     if we dont infer the values for b and c eg a = iota   this prints out the same thing. but this is scoped only to that constant block
// 	b = iota                                                b           if a new constant block is opened, it starts counting from 0.
// 	c = iota                                                c
// )

// func main() {
// 	fmt.Printf("%v\n", a)       this prints out: 0
// 	fmt.Printf("%v\n", b)                        1
// 	fmt.Printf("%v\n", c)                        2

// iota changes values as the constants are being evaluated.

// Example: import (
// 					"fmt"
//						 )

//			const (
// 					catSpecialist = iota
// 					dogSpecialist
// 					snakeSpecialist
//                 )

// 			func main() {
// 					var specialistType int = catSpecialist
// 					fmt.Printf("%v\n", specialistType == catSpecialist)

// when we do not assign a value to specialist type above. the statement will only print true for catSpecialist. this is because catSpecialist has a value 0.
// but if we dont want it that way we can assign iota to errorSpecialist or to  (_) this is called write only variable and the cat specialist will be assigned the value 1 instead.
// this is because iota must be assigned a 0 value.
//                  _ = iota         or       errorSpecialist = iota
//                  catSpecialist
// 					dogSpecialist
// 					snakeSpecialist
// we can also add to the fist value and it will continue in that order.
//                  _ = iota + 5
//                  catSpecialist          (this prints 6)
// 					dogSpecialist          (this prints 7)
// 					snakeSpecialist        (this prints 8)

// we can do bitshifting in place of raising to a particular power.

// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

// func main() {
// 	fileSize := 4000000000.
// 	fmt.Printf("%.2fGB", fileSize/GB)   => this prints 3.73GB this is in 2 d.p because of the .2f

// this whole operation is used to convert the number to GB or MB or KB etc.
//  1 << 0 = 1, 1 << 1 = 2, 1 << 2 = 4, 1 << 3 = 8, 1 << 4 =16

// we can store data into a byte in such a manner.
//package main

// import (
// 	"fmt"
// )

// const (
// 	isAdmin = 1 << iota
// 	isHeadquarters
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmerica
// )

// func main() {
// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
// 	fmt.Printf("%b\n", roles) this prints the value using the OR function on the variable roles.
// 	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin) this will show whether it is isAdmin
// 	fmt.Printf("Is HR? %v", isHeadquarters&roles == isHeadquarters) this will show whether it is isAdmin
// }

// PascalCase => this is when the first letter of each word is capital.
// camelCase => this is when the first letter of the first word is not capital.
