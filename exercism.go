package main

// import (
//     "fmt"
// )
// // TODO: define the 'OvenTime' constant

// const OvenTime int = 40

// // RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
// func RemainingOvenTime(actualMinutesInOven int)  int{
//     actualMinutesInOven = 10
//     return (OvenTime - actualMinutesInOven)
//     }

// // PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
// func PreparationTime(numberOfLayers int) int {
// 	numberOfLayers = 2
//     fmt.Println(numberOfLayers * 2)
//     return 4
// }

// // ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
// func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
// 	numberOfLayers = 3
//     actualMinutesInOven = 20
//     fmt.Println((numberOfLayers * 2) + actualMinutesInOven)
//     return 26
// }

// package techpalace

// import (
//     "strings"
// )

// // WelcomeMessage returns a welcome message for the customer.
// func WelcomeMessage(customer string) string {
// 	return "Welcome to the Tech Palace, "+ strings.ToUpper(customer)
// }

// // AddBorder adds a border to a welcome message.
// func AddBorder(welcomeMsg string, numStarsPerLine int) string {
//     var stars string= strings.Repeat("*",numStarsPerLine)
// 	return stars + "\n" + welcomeMsg + "\n" + stars
// }

// // CleanupMessage cleans up an old marketing message.
// func CleanupMessage(oldMsg string) string {
//     var it1 string = string.Replace(oldMsg, "*", " ", -1)
//     var it2 string = string.TrimSpace(it1)
// 	return it2
// }

// package cards

// // FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
// func FavoriteCards() []int {
//     cards := []int{2, 6, 9}
// 	return cards
// }

// // GetItem retrieves an item from a slice at given position.
// // If the index is out of range, we want it to return -1.
// func GetItem(slice []int, index int) int {
//     // card := GetItem([]int{2, 6, 9}, index)
//     var card int
//     var length int = len(slice) - 1
//     if index > length || index < 0{
//         card = -1
//     } else if index <= length {
//         card = slice [index]
//     }
// 	return card
// }

// // SetItem writes an item to a slice at given position overwriting an existing value.
// // If the index is out of range the value needs to be appended.
// func SetItem(slice []int, index, value int) []int {
//     b := slice
//     var length int = len(slice) - 1
//      if index > length || index < 0{
//         b = append(slice, value)
//     } else if index <= length {
//     	b[index] = value
//     }
// 	return b
// }

// // PrependItems adds an arbitrary number of values at the front of a slice.
// func PrependItems(slice []int, values ...int) []int {
//     c := []int(values)
//     c = append (c, slice...)
// 	return c
// }

// // RemoveItem removes an item from a slice by modifying the existing slice.
// func RemoveItem(slice []int, index int) []int {
//     var length int = len(slice) - 1
//     b := slice
//     if index > length || index < 0 {
//         b = slice
//     } else if index <= length {
//         b = append(slice [:index], slice[index+1:]...)
//     }
// 	return b
// }
