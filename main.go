package main

import "fmt"

// we used interfaces to define a method set or a fcn set
// anything that matches this interface, a getGreeting method returning a type string is now of type bot
// can use these types for any other locations where we expect to see a type of bot

type bot interface {
	getGreeting() string
}

// interfaces are not generic types. other languages have generic types, go does not
// interfaces are implicit, we do not manually have to say that our custom type satisfies some interface
// interfaces are a contract to help us manage types. garbage in > garbage out. if our custom's type implementation of a fcn is broken then interfaces will not help us
// interfaces are tough. step 1 is understanding how to read them. understand how to read interfaces in the standard lib. writing your own interfaces is tough and requires exp

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// no return value as it will receive a value and print it out
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// can omit the value eb and sb if we are not using it and the funcs would still do the same thing
// englishBot and spanishBot are receivers

func (eb englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
