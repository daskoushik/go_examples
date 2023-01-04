package main

import "fmt"

var global string

func main() {

	var a string
	var b string = "Explicit type declaration"
	var c = "Implicit type declaration - type infered by compiler"
	d := "Within a function"
	var i, f = 10, 18.76
	g := 11
	
	global = "I am global variable"
	a = "declaring now"
    fmt.Println("go" + "lang")
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
	fmt.Println("All string variables: ",a, b, c, d, global)
	fmt.Println("Other variables: ",i,f,g)
}