package main
import ("fmt")

func slices(){

	/*
	- similar to arrays in some fashion but have no defined lengths
	- use len() for length and cap() fo capacity
	*/

	var scores = []string{} // this is a slice 0 length and 0 capacity 

	fmt.Print(scores)

	//  make slice frm array
	var myArray = [6]int{56,43,24,22,44}
	var mySlice = myArray[1:3]

	fmt.Println(mySlice)


	// use make()
	/* capacity means no reallocation teill the 10th element, good for memory optimization */
	var slice1 = make([]int, 5, 10)
	var slice3 = make([]int, 1)
	fmt.Print(slice1)

	// add elements to the slice
	slice1 = append(slice1,3,5) // if you want to extend the same slice
	slice3 = append(slice1, mySlice...)
	fmt.Print(slice3)

	// how to copy a slice into a new variable
	var numbers = []int{23,65,33,22,43}
	var numbersCopy = make([]int, len(numbers))
	copy(numbersCopy, numbers)
	numbersCopy[0] = 99
	fmt.Print(numbersCopy)

}