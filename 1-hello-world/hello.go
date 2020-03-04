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

    arr := [5][5]int{{1, 2, 3, 4, 5},
                     {6, 7, 8, 9, 10},
                     {11, 12, 13, 14, 15},
                     {16, 17, 18, 19, 20},
                     {21, 22, 23, 24, 25}}
    fmt.Println("arr = ", arr)
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Print(arr[i][j], " ")
        }
        fmt.Println("")
    }

    // slice 数组
    s1 := make([]string, 3)
    s1[0] = "ko ni chi ha"
    s1[1] = "se kai"
    s1[2] = "!"
    s1 = append(s1, "Hello ")
    s1 = append(s1, "World")
    fmt.Println("s1 = ", s1)
    s2 := make([]string, len(s1))
    copy(s2, s1)
    s2 = append(s2, "!!!")
    fmt.Println("s2 = ", s2)

    // 关联数组
    myMap := make(map[string]string)
    myMap["Name"] = "Harry"
    myMap["Surname"] = "Potter"
    anotherMap := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("myMap = ", myMap)
    fmt.Println("anotherMap = ", anotherMap)

    for key, value := range myMap {
        fmt.Printf("%s -> %s\n", key, value)
    }
}
