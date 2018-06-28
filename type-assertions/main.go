package main

import "fmt"

/* Type Assertions:
 *
 * PrimaryExpression.(Type)
 *
 * - primary expression can be an identifier, element of an array, slice etc.
 * - type can be a type identifier or type literal
 * - assertion is ussed when you are dealing with an Interface, whereas conversion should be used when dealing with Types
 */

type Dog interface {
	bark()
}

type Cat interface {
	meow()
}

type Beagle struct {
	name string
}

func (s Beagle) bark() {
	fmt.Println("Woof! I'm a beagle!")
}

func main() {
	var snoopy Dog
	snoopy = Beagle{"Snoopy"}

	fmt.Println(snoopy)
	isDog, ok := snoopy.(Dog)
	fmt.Println(isDog, ok)
	// Outputs {Snoopy} true
	// snoopy is of type Dog, since it implements
	// bark() which is required for the Dog interface

	isCat, catOk := snoopy.(Cat)
	fmt.Println(isCat, catOk)
	// Outputs <nil> false

	// Panic:
	//
	// Note that if we did not assign the error, then we will get a panic
	//
	// isCatSecondTry := snoopy.(Cat)
	// fmt.Println(isCatSecondTry)
	// --> will give us panic: interface conversion: main.Beagle is not main.Cat: missing method meow
}
