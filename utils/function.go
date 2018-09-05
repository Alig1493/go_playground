package utils 

import (
    "fmt"
)


func plus(a int, b int) int {
    return a + b
}


func double_plus(a, b, c int) int{
    return a + b + c
}


func multiple_returns(a, b, c int) (int, int) {
    return (a + b), (b + c)
}


func variadic_function(nums ...int) {
    fmt.Println(nums, " ")
    
    var total int
    
    for _, num := range nums {
        total += num
    }
    
    fmt.Println("Total: ", total)
}


func closure() func() int {
    
    i := 0
    return func() int {
        i++
        return i
    }

}


func TestFunc() {

    plus := plus(1,2)
    fmt.Println("1+2 =", plus)
    
    double_plus := double_plus(3,4,5)
    fmt.Println("3+4+5 =", double_plus)
    
    val1, val2 := multiple_returns(1, 2, 3)
    fmt.Printf("First value: %d\nSecond value: %d\n", val1, val2)
    
    fmt.Println("Testing variadic functions: ")
    nums := [] int {1,2,3,4}
    nums = nums[:]
    variadic_function(nums ...)
    
    fmt.Println("Testing closure: ")
    value := closure()
    fmt.Println(value())
    fmt.Println(value())
    fmt.Println(value())
    // to start anew inistialize closur again seperately
    // closure returns a function which in turn must be called as well

}

