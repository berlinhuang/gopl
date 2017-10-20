// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:len(os.Args)])
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[:2])
	for i := 1; i < len(os.Args); i++ {//++i  j=i++ are all wrong
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//!-
