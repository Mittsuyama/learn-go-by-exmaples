package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func returnManyValue() (int, int, int, int) {
    return 2, 3, 5, 7
}

func main() {
   fmt.Println(add(1, 2))
}
