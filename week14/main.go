package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 202444099
	student1.name = "박소이"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)

	var student2 student
	student2.id = 202444098
	student2.name = "김지윤"
	student2.gpa = 4.5
	fmt.Println(student2.id)
}
