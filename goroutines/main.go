/* Goroutine:
 * - lightweight execution thread
 * - spawned goroutines depend on the main goroutine
 * - once the main goroutine ends, the others will end as well
 */
package main

import (
	"fmt"
	"time"
)

/* Prints Numbers 0 - 4
 * Waits 100 milliseconds in between
 */
func printNumbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
	fmt.Println("\nexiting print numbers...")
}

/* Prints letters A - E
 * Waits 125 milliseconds in between
 */
func printLetters() {
	for c := 'A'; c <= 'E'; c++ {
		time.Sleep(125 * time.Millisecond)
		fmt.Printf("%c", c)
	}
	fmt.Printf("\nexiting print letters...")
}

/* Goroutines that are running:
 *
 * printNumbers
 * printLetters
 * main goroutine
 *
 * Note that printNumbers and printLetters goroutines
 * are running concurrently.
 *
 * When a new goroutine is started, the goroutine call returns immediately.
 * The control returns to the next line of code, without waiting for the goroutine
 * to finish executing.
 *
 * If the main goroutine terminates, all other goroutines will terminate as well.
 * In this case, we want the main goroutine to run until the showLetters and printNumbers
 * goroutines finish execution.
 *
 *
 * Output from running this program:
 * 0A1B2C3D4
 * exiting print numbers...
 * E
 * exiting print letters...
 * exiting main goroutine...
 */
func main() {
	go printNumbers()
	go printLetters()
	time.Sleep(650 * time.Millisecond)
	fmt.Println("\nexiting main goroutine...")
}
