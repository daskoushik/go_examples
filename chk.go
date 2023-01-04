package main

import "fmt"

func intSeq() func(int)  (int, int) {
    i := 0
    return func(x int) (int, int) {
        i++
        return i, x+1
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt(1))
    fmt.Println(nextInt(2))
    fmt.Println(nextInt(10))

    newInts := intSeq()
    fmt.Println(newInts(11))
}