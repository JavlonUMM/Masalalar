package main

import "fmt"

func main() {
    var son int
    fmt.Print("Uch xonali son kiriting: ")
    fmt.Scan(&son)
    
    birliklar := son % 10
    onliklar := (son / 10) % 10
    
    fmt.Printf("Birliklar xonasidagi raqam: %d\n", birliklar)
    fmt.Printf("O'nliklar xonasidagi raqam: %d\n", onliklar)
}
