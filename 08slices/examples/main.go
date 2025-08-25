package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Id    int
	Score int
}

func main() {
	scores := []int{10, 6, 3, 9, 6, 0, 8, 9}

	scoresAbove5 := []int{}
	scoreBelow5 := []int{}

	for i := range scores {
		if scores[i] >= 5 {
			scoresAbove5 = append(scoresAbove5, scores[i])
		} else {
			scoreBelow5 = append(scoreBelow5, scores[i])
		}
	}

	fmt.Printf("There are %v student that scored 5 and above\n", len(scoresAbove5))
	fmt.Printf("%v scored below 5\n", len(scoreBelow5))

	students := studentWithScore()

	studentWithPass := []Student{}

	for i := range students {
		if students[i].Score >= 5 {
			studentWithPass = append(studentWithPass, students[i])
		}
	}

	for i := range studentWithPass {
		fmt.Printf("%v passed with score of %v\n", studentWithPass[i].Name, studentWithPass[i].Score)
	}

}

func studentWithScore() []Student {
	students := []Student{
		{Name: "Jane Doe", Id: 1, Score: 5},
		{Name: "John Doe", Id: 2, Score: 8},
		{Name: "Jacob barn", Id: 3, Score: 6},
		{Name: "Czar czar", Id: 4, Score: 4},
	}
	return students
}
