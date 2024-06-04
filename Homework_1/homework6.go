package main

import "fmt"

func main() {
    var second int
    fmt.Print("Enter the second: ")
    fmt.Scan(&second)
    
    hour := second / 3600
    
    fmt.Printf("Result: %d\n", hour)
}
