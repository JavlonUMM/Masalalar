package main

import "fmt"

func main() {
    var second int
    fmt.Print("Enter the second: ")
    fmt.Scan(&second)

    minute := second / 60
    
    fmt.Printf("Result: %d\n", minute)
}
