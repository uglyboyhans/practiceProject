package main

import (
	"math"
	"fmt"
)

type Person struct {
	name string
	age  int
}

const (
	WHITE = iota // 用于常量中初始化计数
	BLACK
	RED
	BLUE
	GREEN
)

type Color byte

type Box struct {
	width, heigth, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.depth * b.heigth
}

func (b *Box) setColor(c Color) { // !!!接收者必须是指针？
	b.color = c
}

func (bl BoxList) biggestBoxColor() Color {
	bv := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if v := b.Volume(); v > bv {
			bv = v
			k = b.color
		}
	}
	return k
}

func (bl BoxList) plantThemBlack() {
	for i, _ := range bl {
		bl[i].setColor(BLACK)
	}
}

func (color Color) colorToString() string {
	strs := []string{"white", "black", "red", "blue", "green"}
	return strs[color]
}

func testBoxColors() {
	boxes := BoxList{
		Box{4, 3, 4, RED},
		Box{5, 7, 12, WHITE},
		Box{7, 3, 24, BLUE},
		Box{2, 3.55, 7.25, GREEN}, // 最后一个也要逗号
	}
	biggestColor := boxes.biggestBoxColor()
	fmt.Printf("这里有%v个盒子\n", len(boxes))
	println("最大的盒子是", biggestColor.colorToString(), "色的")
	println("接下来全变成黑色：")
	boxes.plantThemBlack()
	println("现在最大的盒子是", boxes.biggestBoxColor().colorToString(), "色的")
	lastV := boxes[len(boxes)-1].Volume()
	println("最后一个盒子有", lastV, "那么大")
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
	fmt.Printf("%v\n", r1.area())
	fmt.Printf("%v\n", c1.area())
}
