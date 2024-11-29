package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 202444099
	student1.name = "박소이"
	student1.gpa = 4.5

	fmt.Println(student1.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 202444098
	student2.name = "김지윤"
	student2.gpa = 4.5

	fmt.Println(student2.id)
}
