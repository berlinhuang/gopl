// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//for k, v := range lst {
	// }
	for _, arg := range os.Args[1:] { //range will generate a pair value
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)




	var str string = "str"
	var str1 ="str1"
	str2 := "str2" //only used in func
	fmt.Println(str, str1, str2)


}

//!-
