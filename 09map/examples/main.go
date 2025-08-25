package main

type Student struct {
	FullName string
	Email    string
	Id       string
}

type StudentScore struct {
	StudentId  string
	CourseCode string
	Score      float32
}

type Course struct {
	Code string
	Name string
}

func main() {
	students := students()
	courses := courses()

	studentScores := make(map[int]StudentScore)

	var mustapha Student = Student{}

	for i := range students {
		if students[i].Id == "20/05/04/005" {
			mustapha = students[i]
		}
	}

	studentScores[1] = StudentScore{
		StudentId:  mustapha.Id,
		CourseCode: courses[1].Code,
		Score:      76.5,
	}

}

func students() map[int]Student {
	students := make(map[int]Student)

	students[1] = Student{"Ibrahim Mustapha", "iamustapha213@gmail.com", "20/05/04/005"}
	students[2] = Student{"Ibrahim Ibrahim", "lilbro@gmail.com", "20/05/04/001"}

	return students
}

func courses() map[int]Course {
	courses := make(map[int]Course)

	courses[1] = Course{"CPE101", "Intro to programming"}
	courses[2] = Course{"CPE401", "Computer Architecture and Organization"}

	return courses
}
