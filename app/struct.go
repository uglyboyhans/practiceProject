package main

import (
	"math"
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	mark int
	age  int
}

func older(p1, p2 Person) (olderName string, diff int) {
	if p1.age > p2.age {
		return p1.name, p1.age - p2.age
	} else {
		return p2.name, p2.age - p1.age
	}
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() (area float64) {
	area = r.height * r.width
	return area
}

func (c Circle) area() (area float64) {
	area = math.Pi * c.radius * c.radius
	return area
}

func testStruct() {
	r1 := Rect{3, 5}
	c1 := Circle{3.3}
	fmt.Printf("%v\n",r1.area())
	fmt.Printf("%v\n",c1.area())
}
