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
	//1. for  initialization ; condition ; post
	for i := 1; i < len(os.Args); i++ {//++i  j=i++ are all wrong
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)


	var str, i = "",1
	//2. for condition          a traditional "while" loop
	for i<len(os.Args){
		str +=  os.Args[i]
		i++
	}
	fmt.Println(str)


	//3. endless loop              a traditional infinite loop
	for {
		os.Exit(0)
	}

}

//!-
