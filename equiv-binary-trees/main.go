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

import "golang.org/x/tour/tree"
import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		if t.Left != nil {
			Walk(t.Left, ch)
		} else {
			ch <- t.Value
		}

		if t.Right != nil {
			Walk(t.Right, ch)
		}
	}

	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	tree1 := tree.New(1)
	// tree2 := tree.New(2)

	// fmt.Println(Same(tree1, tree2))
	fmt.Println(Same(tree1, tree1))
}
