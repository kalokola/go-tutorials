package main
import ("fmt")

// short for structure)

type Person struct {
	name string
	age int
	weight float64
}

func (person Person) isAdult() bool {
	return person.age > 18
}

func getInfo(){

	// you can populate like this or define an empty struct like initialising variables
	var person Person = Person {
		name: "Bright",
		age: 34,
		weight: 67.8,
	}

	// punch data into the struct
	// person.age = 30
	// person.age = 25

	// age := personData(person)
	fmt.Println(person.isAdult())
}

// you can pass struct as an argument
func personData(person Person) int{
	return person.age
}