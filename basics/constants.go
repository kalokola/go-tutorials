package main
import ("fmt")

func showConstants(){

	/*
	- constants are unchangeable and read only
	- syntax: const VariableName = value
	- replaces var with const

	advice:
	- use upper cases only for variables for easy identification
	*/
	const PI float32 = 3.14 // typed
	const GRAVITY  = "9.18" // inferred
	// setense := fmt.Sprintf("Value of Pi is %v Gravity is %v", PI, GRAVITY)
	// fmt.Println(setense)

	// works same
	fmt.Printf("Value of Pi is %v Gravity is %v\n", PI, GRAVITY)

}