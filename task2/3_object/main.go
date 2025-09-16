package main

import "fmt"

// task3.1
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	a float32
	b float32
}

func (r *Rectangle) Area() {
	fmt.Println("The Rectangle Area =", r.a*r.b)
}

func (r *Rectangle) Perimeter() {
	fmt.Println("The Rectangle Perimeter =", 2*(r.a+r.b))
}

type Circle struct {
	r float32
}

const pi = 3.14

func (c *Circle) Area() {
	fmt.Println("The Cricle Area =", pi*c.r*c.r)
}

func (c *Circle) Perimeter() {
	fmt.Println("The Crycle Perimeter =", 2*pi*c.r)
}

// task3.2
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e *Employee) PrintInfo() {
	fmt.Println("EmployeeID:", e.EmployeeID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
}

func main() {
	// task3.1
	// r := Rectangle{2, 2}
	// r.Area()
	// r.Perimeter()
	// var sr Shape = &r
	// sr.Area()
	// sr.Perimeter()

	// c := Circle{3}
	// c.Area()
	// c.Perimeter()
	// var sc Shape = &c
	// sc.Area()
	// sc.Perimeter()

	// task3.2
	employeeA := Employee{
		1,
		Person{"Jack", 18},
	}
	employeeA.PrintInfo()
}
