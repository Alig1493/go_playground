package utils 

import (
    "fmt"
)

func Range() {

    nums := [] int {2, 3, 4}
    sum := 0
    
    for _, num_value := range nums {
        sum += num_value
    }
    
    fmt.Println("Sum: ", sum)
    
    kvs := map[string]string {"a": "apple", "b": "banana"}
    
    for key, value := range kvs {
        fmt.Printf("%s -> %s\n", key, value)
    }
    
    for i, c := range "go" {
        fmt.Printf("%d - > %c(%d)\n", i, c, c)
    }
    
}

