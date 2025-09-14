package main

// import (
// 	"fmt"
// )

// func main() {
// 	statePopulations := map[string]int{
// 		"California":   39250017,
// 		"Texas":        27862596,
// 		"Florida":      20612438,
// 		"New York":     19745289,
// 		"Pennsylvania": 12802503,
// 		"Illinois":     12801539,
// 		"Ohio":         11614373,
// 	}
// 	for _, v := range statePopulations {
// 		fmt.Println(v)
// 	}
// }

// Basic for loop
// Example:
// func main() {                                      => it starts with a for keyword (expected). then what follows are three statements.
// 	for i := 0; i < 5; i++ {                             the first statement is going to be an initializer, same as initializer in if and switch statements. it is normally used to set up a counter, when using loops.
// 		fmt.Println(i)                                   the second statement is going to be some kind of statement that generates a boolean result, it is used by a for loop, to determine if it is done looping or not.
// 	}                                                    the third statement is the incrementer, so in a situation, it is used to increment a counter variable.
// }
// this prints out 0 1 2 3 4 printed on different lines sequentially.

// func main() {
// 	for i := 0; i < 5; i = i + 2 {
// 		fmt.Println(i)
// 	}
// }                 this prints out 0 2 4 printed on different lines sequentially.

// but the first method is the more common way.

// we have no comma operator in the go language, so we cant seperate 2 statements using comma.
// for i := 0, j := 0; i < 5; i++, j++ {           this is wrong in GO language.

// the correct syntax is to
// func main() {                                                    0   0
// 	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {                  =>  1   1
// 		fmt.Println(i, j)                                           2   2
// 	}                                                               3   3
// }                                                                4   4

// using the incrementer sign for 2 variables, is not possible
// for i, j := 0, 0; i < 5; i, j = i++, j++     this won't print, because the increment operation is not an expression so we cant use it as part of a statement, it is a statement on its own.

// it is very bad practice to manipulate the counter within the for loop. This may cause loops to act really strange.
// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		if i%2 == 0 {
// 			i /= 2                                                   +> this is an example.
// 		} else {
// 			i = 2*i + 1
// 		}
// 	}
// }

// we dont need to use all 3 statements, for example when the variable is initialized at any other point in the application we can leave the first statement out.
// example:
// func main() {
// 	i := 0
// 	for ; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// The difference between the above format and the previous one is that, the above one i is scoped to the main function, while the previous one i is scoped to the for loop.

// another statement we can remove is the third statement.
// example:
// func main() {
// 	i := 0
// 	for i < 5 {
// 		fmt.Println(i)
// 		i++
// 	}
// }                                     => this will still run the same way.

// INFINITE FOR LOOP: Sometimes we need a test for an undefined number of times, especially when we don't know the amount of times we want it to run, this we can create a complex expression that will help generate the time to stop.
// to do such we use a break statement, it is moe commonly used in for loops than in switch statements.
// for example:
// func main() {
// 	i := 0
// 	for {
// 		fmt.Println(i)
// 		i++
// 		if i == 5 {
// 			break
// 		}
// 	}
// }
// when we do this, we leave the entire for loop, and the entire for loop stops.

// we can also use a continue statement.
// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }
// this means that after the if statement is done we would go back to the beginning of the for loop. and continue executing, if the variable i meets the requirements, it them prints out.
// continue statements aren't used very often, but they are useful when you are processing a large set of numbers and you need to determine within the loop whether you want to process the loop or not.

// NESTED LOOP
// func main() {
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			fmt.Println(i * j)
// 		}
// 		fmt.Println(i)
// 	}
// }
// this prints 1 2 3 1 2 4 6 2 3 6 9 3
// this means for all values of the first for loop, we run the second the loop and then print the statement.

// but what if we want to leave the loop as soon as we get something greater than 3;
// example:
// func main() {
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			fmt.Println(i * j)
// 			if i*j >= 3 {
// 				break
// 			}
// 		}
// 	}
// }                                   => this doesnt print the expected result, because the break function only breaks out of the nearest loop. so the first for loop will still continue running.
// to solve this we use a concept called label which we can add.
// func main() {
// 	Loop:
// 		for i := 1; i <= 3; i++ {
// 			for j := 1; j <= 3; j++ {
// 				fmt.Println(i * j)
// 				if i*j >= 3 {
// 					break Loop
// 				}
// 			}
// 		}
// 	}
// you add the label at the beginning of the main function and then you can all it after the break keyword, and this defines where we want to break out.

// HOW TO WORK WITH COLLECTIONS WITH FOR LOOPS.
// example a slice, an arbitrarily sized slice. and maybe you want to be able to loop through all the items in the slice.
// the way we do this is by a special form of the for loop which is called the for range loop.
// func main() {
// 	s := []int{1, 2, 3}
// 	for k, v := range s {
// 		fmt.Println(k, v)
// 	}
//}                                      the variable k stands for key, while the variable v stands for value. so this will show each key and the corresponding value.
// and this is the only syntax we have to remember, when using the for range loop, this key, value syntax, this is what we always get.

// so this works for even arrays, slices, maps, even strings used as the source of the for range loop. this brings each letter. and there key.

// CHANNELS: we can also range over channels. they are used for concurrent programming in GO. we will talk abou later.

// incase we want to see the keys seperate or the values seperate.
// for keys.
// func main() {
// 	statePopulations := map[string]int {
// 		"California": 39250017,
// 		"Texas": 27862596,
// 		"Florida": 20612438,
// 		"New York": 19745289,
// 		"Pennsylvania": 12802503,
// 		"Illinois": 12801539,
// 		"Ohio": 11614373,
// 	}
// 	for k := range statePopulations {                       => this will work out.
// 		fmt.Println(k)
// 	}
// }

// for values. we can use the write only variable.(_)
// func main() {
// 	statePopulations := map[string]int{
// 		"California":   39250017,
// 		"Texas":        27862596,
// 		"Florida":      20612438,
// 		"New York":     19745289,
// 		"Pennsylvania": 12802503,
// 		"Illinois":     12801539,
// 		"Ohio":         11614373,
// 	}
// 	for _, v := range statePopulations {
// 		fmt.Println(v)
// 	}
// }
