package main
import ("fmt")

func greetigMessage(message string){
	fmt.Println(message)
}

func totalScore(score1 int, score2 int) int {
	/* the return type must be specified */
	return score1 + score2
}

func totalMarks(score1 int, score2 int) (marks int ){
	/* the return type must be specified */
	marks = score1 + score2
	return 
}