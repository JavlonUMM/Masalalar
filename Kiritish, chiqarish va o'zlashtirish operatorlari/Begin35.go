package main

import "fmt"

func main() {
    var V, U, T1, T2 float64
    fmt.Print("Qayiqning tezligini kiriting: ")
    fmt.Scan(&V)
    fmt.Print("Daryo oqimini tezligini kiriting: ")
    fmt.Scan(&U)
    fmt.Print("Daryo oqimiga qarshi harakat vaqti kiriting: ")
    fmt.Scan(&T2)
    fmt.Print("Daryo oqimi boyicha harakat vaqti kiriting: ")
    fmt.Scan(&T1)

    oqimga_qarshi := T2 * V

    oqimi_boyicha_yol := T1 * (V - U)

    umumiy := oqimga_qarshi + oqimi_boyicha_yol

    fmt.Printf("Qayiq %2f km yo'l bosdi\n", umumiy)
}
