package main

import "fmt"

// Parameterized Types
type Radius interface {
    int64 | int8 | float64
}
 
func generic_circumference[R Radius](radius R){ 
    var c R
    c = 2 * 3 * radius
    fmt.Println("The circumference is: ", c)
}
