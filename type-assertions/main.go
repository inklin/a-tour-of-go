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
}
