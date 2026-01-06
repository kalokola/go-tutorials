package main
import ("fmt")

var firstName string = "Trixa" // typed and global use var
func variablesInfo(){
	/*
	- int
	- float32
	- string
	- bool
	syntax: var variableName type = value // for global scopes
	or just variableName := value // local spaces
	*/

	// * variables can be declared without values
	var age int

	fmt.Println(age)

	age = 30 // you can assign values like this
	fmt.Println(age)

	// * type can be inferred or typed var firstName string 
	var lastName = "Limited" // inferred

	// short notation for infered
	notation := "LLC"

	fmt.Printf("My name is %s %s - %s",firstName, lastName, notation)


	var fName,lName,tld string = "Zipa","Africa",".tz" // multiple assignments
	language, country := "sw","TZ"

	fmt.Printf("\n%s %s", language,country)
	fmt.Printf("\n%s %s %s\n", fName,lName,tld)

	// block declaration of variables
	var (
		a = 45
		b = 54
		c = 12.5
	)

	fmt.Print(a , b , c)


	/*
	rules for naming varibles
	- must start with letter or under score _
	- can't start with a digit
	- can only contain alpha numeric characters like a-z, A-Z, 0-9 and _
	- variables are case sensitive eg age and Age or AGE agE
	- no limit of length of the vvariable
	- no spaces
	- avoid go key words
	*/


	// camel eg firstName, pascal eg FirstName and snake eg first_name are allowed in Go
}
