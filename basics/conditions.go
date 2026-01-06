package main
import ("fmt")

func conditions(){
	/*
	- if
	- if else
	- if else if
	- switch
	*/
	var age int = 4
	var threshold int = 18
	if age > threshold { // where 20 > 6 is the condition
		fmt.Println("Adult")
	} else if (age > 4 && age < 7){
		fmt.Println("Toddler")
	} else {
		fmt.Println("Infant")
	}

	

	ageSwitch(5,18)
	scoreSwitch(40)
	
}

func ageSwitch(age int, threshold int){
		switch  {
		case age > threshold:
			fmt.Print("Adult")
		case age > 4 && age < 7:
			fmt.Println("Toddler")
		default:
			fmt.Println("Infant")
		
	}
}

func scoreSwitch(score float64){
		switch  score {
		case 39,39.5: // multiple switches just like chcking if i
			fmt.Print("Would Sup")
		default:
			fmt.Println("Infant")
		
	}
}