package utils

import (
    "fmt"
)

func Map() {
    // Declare a map with string keys and integer values
    m := make(map[string] int)
    
    m["k1"] = 7
    m["k2"] = 13
    
    fmt.Println("map: ", m)
    
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
    fmt.Println("len: ", len(m))
    
    delete(m, "k2")
    fmt.Println("map: ", m)
    
    for k := range m {
        value, prs := m[k]
        fmt.Println("value", value)
        fmt.Println("prs: ", prs)
    }
    
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map: ", n)
    
}

