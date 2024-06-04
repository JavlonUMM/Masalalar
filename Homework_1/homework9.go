package main

import "fmt"

func main() {
    var number int
    fmt.Print("Enter a number greater than 999: ")
    fmt.Scan(&number)
    
    result := (number / 1000) % 10
    
    fmt.Printf("Result: %d\n", result)
}
