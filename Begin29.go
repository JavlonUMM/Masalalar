package main

import "fmt"

func main() {
    var alfa float64
    fmt.Print("Gradusda berilgan burchakni kiriting (0° < alfa < 360°): ")
    fmt.Scan(&alfa)
	const pi = 3.14
    
    radian := alfa * (pi / 180)
    
    // Natijani chiqarish
    fmt.Printf("Berilgan burchakning qiymati radianda: %.4f\n", radian)
}
