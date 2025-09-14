package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Animal struct {
// 	Name   string 'required max: "100"'
// 	Origin string
// }

// func main() {
// 	t := reflect.TypeOf(Animal{})
// 	field,_:= t.FieldByName("Name")
// 	fmt.Println(field.Tag)
// }

// What are maps: this take a key(state names) and maps it to some kind of a value(population), it provides a flexible datatype.
//example: map of Us state names, with their population
// func main() {
// 	statePopulations := map[string]int{                string => keytype, int => valuetype
// 		"California":   39250017,
// 		"Texas":        27862596,
// 		"Florida":      20612439,
// 		"New York":     19745289,
// 		"Pennsylvania": 12802503,
// 		"Illinois":     12801539,
// 		"Ohio":         11614373,
// 	}
// 	fmt.Println(statePopulations)
// }

// we can use different type for the keys and any type for the values, but when we declare a map that type has to be consistent for every key value pair within the map.
// we have a lot of options with the keytype, but we dont have an infinite number of options.
// the basic constaints for the keys is that they have to be able to test for equality. All numeric types, strings, pointers, interfaces, arrays, channels can be tested for equality.
// some datatypes that cannot be used for equivalency checking are slices, maps and other functions. a slice cannot be an index to a map. but arrays can.

// ways of creating maps.
// 1. literal syntax: the most common way, map[keytype]valuetype{}
// 	statePopulations := map[string]int{                string => keytype, int => valuetype
// 		"California":   39250017,
// 		"Texas":        27862596,
// 		"Florida":      20612439,
// 		"New York":     19745289,
// 		"Pennsylvania": 12802503,
// 		"Illinois":     12801539,
// 		"Ohio":         11614373,
// 	}
// 	fmt.Println(statePopulations)
// }

// 2. Built in make function: variable := make(map[keytype]valuetype)

// How to manipulate the values of our maps
// to pull out one single value from the map for example the map above:
// fmt.Println(statePopulations["Ohio"]) this is to bring out ohio.

// to add a value to the map.
// for the example above:
// statePopulations["Georgia"] = 10310371. this has been added to the map and can be pulled out anyhow.

// The return order of a map is not predictable, unlike a slice or an array which will return it in the order which we made it.
// in a map, everything is stored in a way that we cannot guarantee the order.

// how to delete entries from a map.
// 	we use the built-in "delete" function
// delete (name of map, name of key you are deleting)

// in a map, if you print out a key value which is not present in the map, it prints out 0. which might be a problem, because this may mean the city typed has 0 population.
// to prevent such an issue, we use (,ok) syntax
// for example:
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
// 	pop, ok := statePopulations["Oho"]
// 	fmt.Println(pop, ok)
// }

// we can check the length of a map;
// by: fmt.Println(len{statePopulations})

// if you start passing maps around, and manipulating it, it will affect the underlying map. manipulating the map in one place is going to have effect on every other place it is used.

//STRUCTS
// this is a list of field names and their types associated with them
// The struct gathers information together that are related to one concept and it does it in a very flexible way.
// we dont have any constraints on the types of data that can be put in structs, we can mix anytype of data together. all the other collection types have had consistent data types.
// structs are the simplest types of collection, they are also some of the most powerful

// this is the field name syntax.
// package main

// import (
// 	"fmt"
// )

// type Doctor struct {
// 	number     int
// 	actorName  string
// 	companions []string
// }

// func main() {
// 	aDoctor := Doctor{
// 		number:    3,
// 		actorName: "Jon Pertwee",
// 		companions: []string{
// 			"Liz Shaw",
// 			"Jo Grant",
// 			"Sarah Jane Smith",
// 		},
// 	}
// 	fmt.Println(aDoctor)
// }

// how to interrogate one value from the struct
// we use the dot syntax
// example: fmt.Println(aDoctor.actorName)  => prints out "Jon Pertwee"
// fmt,Println(aDoctor.companions[1])    => prints out "Jo Grant"

// another way is using the positional syntax to make the struct.
// func main() {
// 	aDoctor := Doctor{
// 		3,
// 		"Jon Pertwee",
// 		[]string{
// 			"Liz Shaw",
// 			"Jo Grant",
// 			"Sarah Jane Smith",
// 		},
// 	}
// 	fmt.Println(aDoctor)         => this prints the same thing
// }
// this is valid, but it is encouraged not to be used, because it can become a maintenance. this is because if there are more values in the struct than the initializer this could cause some problems.
// the positional syntax also requires that the values be declared in order.

// using the field name syntax they dont need to be in the exact order.

// Naming Conventions. Rules
// starts with Cpital letter : it is going to be exported from the package.
// starts with lower case letters: going to be internal.
// to make the field names to be seen by other packages, you make their first letters capital letter. eg instead of "number" you write "Number"
// no underscores and stuff like that.

// another way of creating a struct is called an anonymous struct.

// func main() {
// 	aDoctor := struct{ name string }{name: "John Pertwee"}
// 	fmt.Println(aDoctor)
// }
// the first set of curly braces is paired to the struct and it is defining the structure of the struct
// the second is the initializer and it is providing data into the struct.

// you cant use it anywhere else because it is anonymous, and it has no name to use to call it.
// it is used in relaatively few situations and also short lived.
// it is used in places where you need to structure data in a way that you dont have in a formal type.

// structs are value types.
// func main() {
// 	aDoctor := struct{ name string }{name: "John Pertwee"}
// 	anotherDoctor := aDoctor
// 	anotherDoctor.name = "Tom Baker"
// 	fmt.Println(aDoctor)             => prints {John Pertwee}
// 	fmt.Println(anotherDoctor)       => prints {Tom Baker}
// }

// if you pass a sturct in your application, you are passing copies, they are referring to independent data sets, you are creating a copy anytime you are manipulating this.

// if we want to point to the same underlying data we use the pointer (&) function.

// EMBEDDING
// Go doesnt support trditional object oriented principles eg inheritance.
// instead it uses composition. instead of esablishing the "is a" relationship.
// it supports composition through embedding.
// here we embed the animal struct in the bird struct.
// package main

// import (
// 	"fmt"
// )

// type Animal struct {
// 	Name   string              => fields
// 	Origin string
// }

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly   bool
// }

// func main() {
// 	  b := Bird{}
//    b.Name = "Emu"
//    b.Origin = "Australia"
//    b.SpeedKPH = 48
//    b.CanFly = false
//    fmt.Println(b)
// }

// instead of inheritance saying a bird is an animal, by embedding (composition relationship) you say a bird has an animal or has animal-like characteristics.
// here a bird is not the same thing as an animal, they cannot be used interchangeably.

// using the literal syntax: you have to be aware that you are using embedding, but if i just declare the object and manipulate it from the outside(as done above), you dont have to be aware of embedding.
// Example:
// func main() {
// b := Bird{
// 	Animal:   Animal{Name: "Emu", Origin: "Australia"},
// 	SpeedKPH: 48,
// 	CanFly:   false,
// }
// when should you use embedding.
// Generally for modelling behaviour, embedding is not the right choice. generally it is much better to use interfaces.
// when authoring a library and you are creating a web framework it can be good.

// Tags
// type Animal struct {
// 	Name   string 'required max: "100"'    => this makes the name field required and a max input size of 100
// 	Origin string
// }

// func main() {
// 	t := reflect.TypeOf(Animal{})
// 	field,_:= t.FieldByName("Name")
// 	fmt.Println(field.Tag)                => prints out required max: "100"
// }

// this just shows when maybe a space of strings is required to be filled, and it maximum amount it 100.
// they are added to struct fields to describe that field in some way.

// take some time to read on tags.
