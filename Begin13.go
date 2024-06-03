package main

import "fmt"

func main() {
    var R1, R2 float64
    
    fmt.Print("Birinchi aylananing radiusini kiriting (R1): ")
    fmt.Scan(&R1)
    
    fmt.Print("R1 dan katta ikkinchi aylananing radiusini kiriting (R2): ")
    fmt.Scan(&R2)

	const pi = 3.14
	
    S1 := pi * R1
    S2 := pi * R2
    S3 := pi * (R1 - R2)
    
    fmt.Printf("Birinchi aylananing yuzasi (S1): %.2f\n", S1)
    fmt.Printf("Ikkinchi aylananing yuzasi (S2): %.2f\n", S2)
    fmt.Printf("Aylanalar orasidagi yuzalar farqi (S3): %.2f\n", S3)
}
