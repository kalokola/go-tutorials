package main

import "fmt"

// import "fmt"


func maps(){
	// interface allows differete data types
	var fruits = map[string]interface{}{"name":"Melon", "weight": 0.2}
	salads := fruits /* Maps Are References :: always use copy to avoid mutating maps that share values when eany is updated */

	salads["age"] = 33
	for index, value := range fruits {
		fmt.Println(index,value)
	}

	delete(fruits, "weight") // map, key
	fmt.Println(fruits["weight"]) // comes out as nil

	var infos = make(map[string]int) // initialise empty one
	/* without make the map / any variable is nil, make creates empty */ 
	fmt.Println(infos)

	var mapNil map[string] string
	fmt.Println(mapNil) // empty map[]

	value, exists := fruits["age"] // check if item exists

	fmt.Println(value, exists)


}

