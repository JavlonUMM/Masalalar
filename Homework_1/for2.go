package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a number from 1 to n: ")
    fmt.Scan(&n)
    
    fmt.Println("Result:")
    for i := n; i >= 1; i-=1 {
        fmt.Println(i)
        
    }
}
