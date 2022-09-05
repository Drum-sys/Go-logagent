package main

import "fmt"

type Stu struct {
	Name string
	Age int
}

func (p *Stu) SetName(name string) *Stu {
	p.Name = name
	return p
}

func (p *Stu) SetAge(age int) *Stu {
	p.Age = age
	return p
}

func (p *Stu) Print() *Stu {
	fmt.Printf("name:%s, age:%d\n", p.Name, p.Age)
	return p
}

func main()  {
	stu := &Stu{}
	stu.SetName("ljw").SetAge(12).Print()
	//stu.SetAge(12)
	//stu.Print()
}
