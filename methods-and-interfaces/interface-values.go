package main

import "fmt"

type Character struct {
  name string
  age int
}

type Wizard struct {
  Character
  powerLevel int
}

type Human struct {
  Character
}

type Player interface {
  growOlder()
  selfIntro()
}

// All characters have a self intro function
func (c Character) selfIntro() {
  fmt.Println("My name is", c.name, "and my age:", c.age)
}

// Wizard has a growOlder function
// 
// Note that since this is a non-pointer reference,
// we cannot change the value, we can only read
func (w Wizard) growOlder() {
  fmt.Println("Wizards don't age, I will be forever", w.age, "because of power level:", w.powerLevel)
}

// Human has a growOlder function
// mutates the age, increasing the age by 1
func (h *Human) growOlder() {
  h.age++
}

func main() {
  var merlin Player
  var bob Player

  merlin = Wizard{Character{"Merlin", 100}, 1000}

	// Calling a method on the interface Player
	// executes the method of the same name on the underlying type.
	merlin.selfIntro()
	// In this case, the underlying type is Wizard, and the growOlder function
	// on type Wizard is called.
	merlin.growOlder()
	// When we call selfIntro again, we see that merlin's age is unchanged
  merlin.selfIntro()

  // Note that Human needs to implement
  // a pointer receiver to fulfill the interface of Player
  // that needs a growOlder() function
  bob = &Human{Character{"Bob", 30}}

	// Initially Bob has an age of 30, which we initialized.
	bob.selfIntro()
	// The underlying type is Human, and the growOlder function on type Human is called.
	// Bob's age is mutated and increased by 1. 
	bob.growOlder()
	// When we call selfIntro again, we see that bob's age has increased to 31.
  bob.selfIntro()
}

// Output from this program:
//
// My name is Merlin and my age: 100
// Wizards don't age, I will be forever 100 because of power level: 1000
// My name is Merlin and my age: 100
// My name is Bob and my age: 30
// My name is Bob and my age: 31