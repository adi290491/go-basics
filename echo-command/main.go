// Echo prints its command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var global *int

var a = b + c
var b = f1()
var c = 1

func f1() int {
	return c + 1
}

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Print()
	}

	f()

	*global++
	fmt.Println("Address held by global: ", global)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}

func f() {
	var x int
	x = 1
	global = &x
	fmt.Println("Address of x: ", &x)
}
