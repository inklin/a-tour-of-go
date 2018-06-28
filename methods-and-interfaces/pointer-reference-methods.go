package main

/* Pointer VS. Non-Pointer receiver methods 
 * 
 * Why pass a pointer ? (Why would you want to pass by reference instead of value)
 * - you want to modify the underlying values
 * - the struct is large and deep copying would be expensive
 * - for consistency: if some methods on the struct have pointer receivers, others
 *   should have pointer receivers as well for consistency
 *
 */

import "fmt"

type Coordinate struct {
  x int
  y int
}

// Pass by value, non-pointer receiver
// this will not change the value of x and y
func (c Coordinate) shouldNotMutate() {
  c.x = 100
  c.y = 200
}

// Pointer receiver
// This will allow you to access and change
// the underlying value of the pointer
func (c *Coordinate) shouldMutate() {
  c.x = 1000
  c.y = 1000
}

func main() {
  coordinate := Coordinate{25, 35}
  fmt.Println("Original values of x,y: ", coordinate.x, coordinate.y)

  coordinate.shouldNotMutate()
  fmt.Println("After calling shouldNotMutate():", coordinate.x, coordinate.y)

  coordinate.shouldMutate()
  fmt.Println("After calling shouldMutate():", coordinate.x, coordinate.y)  
}