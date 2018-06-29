/* Channels
 *
 * - Channels are how goroutines communicate with eachother.
 * - can be made with make(chan int)
 * - a typed conduit through which you can send and receive values
 *   using the channel operator <-
 */
package main

import "fmt"

func sum(numbers []int, c chan int) {
	total := 0

	for _, v := range numbers {
		total += v
	}

	// Send the total to the channel
	c <- total
}

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 5, 5, 5}

	// make a channel
	c := make(chan int)

	go sum(s1, c)
	go sum(s2, c)

	// receive the sum results from the channel
	sum1, sum2 := <-c, <-c
	fmt.Println(sum1, sum2)
}
