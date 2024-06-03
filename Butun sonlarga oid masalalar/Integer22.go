package main

import "fmt"

func main() {
    var sekund int

    fmt.Print("Sekundni kiriting: ")
    fmt.Scan(&sekund)

    soatlar := sekund / 3600
    qolgan_sekund := sekund % 3600

    fmt.Printf("O'tgan soatlar va sekundlar: %d soat %d sekund\n", soatlar, qolgan_sekund)
}
