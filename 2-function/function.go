package main

import "fmt"

func Fibonacci(a int) int {
    if a < 3 {
        return 1
    }
    return Fibonacci(a - 1) + Fibonacci(a - 2)
}

func add(a, b int) int {
    return a + b
}

func returnManyValue() (int, int, int, int) {
    return 2, 3, 5, 7
}

func sum(nums ...int) int {
    total := 0
    for _, item := range nums {
       total += item
    }
    return total
}

func newInt() func() int {
    i := 1443
    return func() int {
        i += 451
        return i
    }
}

func main() {
   fmt.Println(add(1, 2))
   fmt.Println(sum(1, 2, 3, 4))
   fmt.Println(sum(1, 2, 3, 4, 5, 6))

   nextInt := newInt()
   fmt.Println(nextInt())
   fmt.Println(nextInt())
   fmt.Println(newInt())

   anotherInt := newInt()
   fmt.Println(anotherInt())
}
