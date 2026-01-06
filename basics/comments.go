package main
import ("fmt")


func comments(){
	// comments are usually ignored by compilers
	// a function to print a full name of brightius kalokola
	/* this is a mutiline comment 
	you can adjusts and place more contents here
	*/
	firstName := "Brightius";
	lastName := "Kalokola";
	_, words := fmt.Printf("My name is  %v %v", firstName, lastName)
	fmt.Println(words)
}