// Created by: haokai
// Created on: May 2021
// This program displays, "Area-of-a-Trapezoid"

package main

import "fmt"

func main() {
	// This function does addition
	var aBase int
	var bBase int
	var height int
	var area int
	// input
	fmt.Println("This program finds the area of a Trapezoid.")
	fmt.Println()
	fmt.Print("Enter the a Base (cm): ")
	fmt.Scanln(&aBase)
	fmt.Print("Enter the b Base (cm): ")
	fmt.Scanln(&bBase)
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)
	fmt.Println()
	// process
	area = (aBase + bBase) / 2 * height
	// output
	fmt.Println("The area is: ", area, " cm².")
	fmt.Println("\nDone.")
}
