package main

import "fmt"

func main() {
    var bayt int
    fmt.Print("Baytni kiriting: ")
    fmt.Scan(&bayt)
    
    kilobayt := bayt / 1024
    
    fmt.Printf("Kilobaytlarda: %d\n", kilobayt)
}
