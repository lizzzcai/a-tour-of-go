/* Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
 */

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// init the map
	wc := make(map[string]int)
	// splict the word from string
	words := strings.Fields(s)
	// iterate through words
	for _, w := range words {
		// get the value and check if exist
		if count, ok := wc[w]; ok == false {
			wc[w] = 1
		} else {
			wc[w] = count + 1
		}
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
