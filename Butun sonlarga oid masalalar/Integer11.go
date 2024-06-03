package main

import "fmt"

func main() {
    var son int
    fmt.Print("Uch xonali son kiriting: ")
    fmt.Scan(&son)
    
    birliklar := son % 10
    onliklar := (son / 10) % 10
    yuzliklar := son / 100
    
    yigindi := birliklar + onliklar + yuzliklar
    
    fmt.Printf("Raqamlar yig'indisi: %d\n", yigindi)
}
