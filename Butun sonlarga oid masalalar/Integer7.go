package main

import "fmt"

func main() {
    var son int
    fmt.Print("Ikki xonali son kiriting: ")
    fmt.Scan(&son)
    
    birliklar := son % 10
    onliklar := son / 10
    
    yigindi := birliklar + onliklar
    
    fmt.Printf("Raqamlar yig'indisi: %d\n", yigindi)
}
