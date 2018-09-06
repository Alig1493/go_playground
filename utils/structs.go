package utils

import (
    "fmt"
)


type People struct {
    name string
    age int
}


func Struct() {
    
    person := People{"Bob", 20}
    fmt.Println(person)
    
    person = People{name: "Alice", age: 30}
    fmt.Println(person)
    
    person = People{name: "Fred"}
    fmt.Println(person)
    
    person = People{name: "Ann", age: 40}
    fmt.Println(&person)
    
    person = People{name: "Sean", age: 50}
    fmt.Println(person.name)
    
    person.age = 100
    fmt.Println(person)

}

