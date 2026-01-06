package main
import ("fmt")

func loops (){
	for i := 0; i <= 10; i++ {
		
		if i == 3 {
			continue // skips excution of the below code and goes to the next iteration
		}

		if i == 6 {
			break // stops excution
		}

		fmt.Println("Counter ", i)
	}

	loopFor()
}

func loopFor(){

	var myArray = []int{23,45,3,22,53}
	// for index, value := range array | slice | map
	for index, value := range myArray {
		fmt.Printf("Index %v %v\n", index, value)
	}
}
