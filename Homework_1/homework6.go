package main

import "fmt"

func main() {
    var second int

    fmt.Print("Enter the second: ")
    fmt.Scan(&second)

    minutes := second / 60
    seconds := second % 60

    fmt.Printf("Result: %d %d\n", minutes, seconds)
}
