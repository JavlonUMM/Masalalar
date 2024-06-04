package main

import (
	"fmt"
)

func main() {
	var Y string 
	var K string 

	fmt.Print("Enter robot's starting direction (north, south, east, west): ")
	fmt.Scan(&Y)

	fmt.Print("Enter moves with 0, 1, 2: ")
	fmt.Scan(&K)

	for move {
		

		switch move {
		case '0':
		case '1':
			switch Y {
			case "s":
				Y = "q"
			case "j":
				Y = "g"
			case "q":
				Y = "s"
			case "g":
				Y = "j"
			}
		case '2':
			
			switch Y {
			case "s":
				Y = "g"
			case "j":
				Y = "q"
			case "q":
				Y = "j"
			case "g":
				Y = "s"
			}
		default:
			fmt.Println("Invalid move:", move)
		}
	}

	fmt.Println("Robot's final direction:", Y)
}
