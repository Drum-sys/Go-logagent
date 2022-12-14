package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}

func main() {
	var root *Student = new(Student)

	root.Name = "stu01"
	root.Age = 12
	root.Score = 98.54

	var left1 *Student = new(Student)
	left1.Name = "stu02"
	left1.Age = 18
	left1.Score = 93

	var left2 *Student = new(Student)
	left2.Name = "stu03"
	left2.Age = 15
	left2.Score = 97

	var right1 *Student = new(Student)
	right1.Name = "stu04"
	right1.Age = 12
	right1.Score = 87

	root.left = left1
	root.right = right1
	left1.left = left2

	trans(root)


}