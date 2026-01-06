package main
import ("fmt")

func dataTypes(){

	// bool
	var isActive bool = true

	fmt.Printf("is logged in ? =  %v\n", isActive)


	// integers can be signed or unsigned
	var age uint = 26 // can only store non negative 
	var degrees int = -26 // can store any value
	fmt.Printf("Age %v at %v", age, degrees)

	var weight float64 = 67.5
	fmt.Print("A melon grew %v", weight)
}