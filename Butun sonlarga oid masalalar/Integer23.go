package main

import "fmt"

func main() {
	var sekund int

	fmt.Print("Sekundni kiriting: ")
	fmt.Scan(&sekund)

	soat := sekund / 3600
	minut := (sekund % 3600) / 60
	qolgan_sekund := (sekund % 3600) % 60

	fmt.Printf("O'tgan soat, minut' sekund: %d %d %d \n", soat, minut, qolgan_sekund)
}
