package utils

import (
    "fmt"
    "time"
)

func Switch() {

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
    
    t := time.Now()
    
    switch t.Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("It's the weekend")
        default:
            fmt.Println("It's a weekday")
            fmt.Println(time.Now())
    }
    
    t = t.Add(time.Duration(-12) * time.Hour)
    
    switch {
        case t.Hour() < 12:
            fmt.Println("It's before Noon")
        default:
            fmt.Println("It's after Noon.")
    }
    
    WhatAmI := func(i interface{}) {
    
        switch t := i.(type) {
            case bool:
                fmt.Println("I'm a bool")
            case int:
                fmt.Println("I'm an integer")
            default:
                fmt.Printf("Unknown Type %T\n", t)
        }
    }
    
    WhatAmI(true)
    WhatAmI(1)
    WhatAmI("hey")
    
}

