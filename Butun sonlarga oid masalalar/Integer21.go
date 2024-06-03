package main

import "fmt"

func main() {
    var sekund int

    fmt.Print("Sekundni kiriting: ")
    fmt.Scan(&sekund)

    qolgan_minut := sekund / 60
    qolgan_sekund := sekund % 60

    fmt.Printf("O'tgan minutlar va sekundlar: %d %d\n", qolgan_minut, qolgan_sekund)
}
