package main

import "fmt"

func main() {

    res := func (a int, b int) int {
    return a + b
	}
    fmt.Println("1+2 =", res(1,2))
	
	func(a int, b int) {
	fmt.Println("Annonymous function multiply: ",a * b)
	}(10, 20)
}