package main

import "fmt"

func main() {

	sugar := 2
	sugar = 4
	// Anonymous function
	coffeeBanao := func() {
		fmt.Println("Making coffee...")
	}
	coffeeBanao()

	// IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Printf("Making %d tea...\n", sugar)
	}()

	// IIFE with parameters
	func(beverage string) {
		fmt.Printf("Making %s...\n", beverage)
	}("juice")

	makeChai := func() {
		chai := "Black tea"
		fmt.Printf("Making... %s with %d spoons of sugar\n", chai, sugar)
	}
	makeChai()
}
