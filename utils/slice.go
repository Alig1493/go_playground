package utils

import (
    "fmt"
)

func Slice() {

    s := make([] string, 3)
    fmt.Println("emp: ", s)
    
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set: ", s)
    fmt.Println("get value of index 2: ", s[2])
    
    fmt.Println("len: ", len(s))
    
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("appended: ", s)
    
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("Copy: ", c)
    
    fmt.Println("sl1 values between index 2 and 5 exclusive: ", s[2:5])
    fmt.Println("sl2 all values less than index 5: ", s[:5])
    fmt.Println("sl3 values from index 2 onwards: ", s[2:])
    
    twoD := make([][] int, 3)
    for i:=0; i<3; i++ {
        inner_length := i + 1
        twoD[i] = make([]int, inner_length)
        for j:=0; j<inner_length; j++ {
            twoD[i][j] = i + j
        }
    }
    
    fmt.Println("2D: ", twoD)

}

