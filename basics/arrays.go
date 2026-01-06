package main
import ("fmt")

func arrays(){
	// arrays are used to store multiple values of the same type in a single variable
	// syntax
	// var variableName = [length]datatype{value1, value2, value3 ....}

	
/*
notes:
	you can infer length as var names = [...]string{"Bright","Kalokola", }
*/

var scores = [3]int{45,89,30}
fmt.Println(scores)

// go is a zero indexing language
fmt.Println(scores[0]) // gives the first element
fmt.Println(scores[1]) // gives the second element

//  arrays are mutable (values can be changed)
scores[0] = 90 // mutate i.e assign a new value
fmt.Println(scores[0]) // gives the first updated element

//  arrays can be fully and partially initilaised i.e they accept numbers less than the length defined but not more
var marks1 = [3]int{}
var marks2 = [3]int{2,6}
fmt.Println(marks1, marks2) //  the above marks1 and marks2 is totally fine


// you can initialise speciifc elements only

var records = [...]int{0:30,2:11}
// arrays dont allow append
fmt.Println(records)
//  use len() to get the length of members in arrays
fmt.Printf("There are %v records i.e %v\n", len(records), records)
}