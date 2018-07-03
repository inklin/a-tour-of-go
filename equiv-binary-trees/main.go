/* Exercise: Equivalent binary trees
 *
 * Determine whether two binary trees have the same sequence of values stored in the leaves.
 * - Use the tree package, the tree type is defined as below
 *
 * type Tree struct {
 *  Left *Tree
 *  Right *Tree
 *  Value int
 * }
 *
 */
package main

import (
	"fmt",
	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
	// reached end of tree
	if t == nil {
		return
	}
	// walk to the very left node of the tree
	walkImpl(t.Left, ch)
	// send the value of the tree node to the channel
	ch <- t.Value
	// walk to the right
	walkImpl(t.Right, ch)
}

// Walk function:
// walks along the tree t and sends all the values
// of the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// Need to close the channel here
	close(ch)
}

// Same determines whether trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Create two channels for the two trees
	ch1, ch2 := make(chan int), make(chan int)

	// Have to goroutines that walk the
	// trees t1 and t2, give them two different
	// channels to receive the node values
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// For every value received in ch1,
	// compare it to value received in ch2
	// to see if the tree values are the same.
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Print("Same tree: tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}

	fmt.Print("Different trees: tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
