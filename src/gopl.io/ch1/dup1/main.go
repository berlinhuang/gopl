// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)// declare a map type counts which keep the pair(key,value)
	input := bufio.NewScanner(os.Stdin) // declare a bufio.NewScanner Type
	for input.Scan() {  //read a text   (input ctl+d to end reading)
		//line := input.Text()
		//counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			//log.Printf   fmt.Errorf         use fmt.Printf
			//function name end with ln(line)    use Println
		}
	}
}

//!-
