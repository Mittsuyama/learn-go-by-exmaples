package main

import "fmt"

type tell interface {
    speak() string
}

type Person struct {
    name string
    gender, age int
}

type Dog struct {
    name, kind string
}

func zeroIt(vptr *int) {
    *vptr = 0
}

func (p Person) getAge() int {
    return p.age
}

func (p Person) speak() string {
    return "Human"
}

func (D Dog) speak() string {
    return "Dog"
}

func tellMeYourSpecies(t tell) {
    fmt.Println("Hello, I am a/an", t.speak())
}

func main() {
    value := 1
    fmt.Println("Origin = ", value)
    zeroIt(&value)
    fmt.Println("Next = ", value)

    peter := Person{"Peter", 1, 18}
    fmt.Println("Peter", peter)
    personPoint := &peter
    personPoint.age = 14
    fmt.Println("young Peter", peter)
    fmt.Println("Peter's age is", peter.getAge())
    samoyed := Dog{"Xiao Sa", "Samoyed"}
    tellMeYourSpecies(peter)
    tellMeYourSpecies(samoyed)
}
