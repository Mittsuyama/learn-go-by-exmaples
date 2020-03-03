package main

import "fmt"

const INF = 0x7fffffff

func main() {
    fmt.Println("Hello World!")
    firstVariable := 1
    fmt.Println("firstVariable =", firstVariable)
    fmt.Println("INF =", INF)

    sum := 0
    for index := 0; index < 101; index ++ {
        sum += index
    }
    // also another method to use for as while in c++
    // count := 0
    // for count < 101 {
    //     sum += count
    // }
    fmt.Println("sum =", sum)

    if sum > 100 {
        fmt.Println("Oh, it's so big!")
    }
}
