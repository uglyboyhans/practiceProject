package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Kodomo struct {
	Human // 匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) sayHi() {
	fmt.Println("Hello,my name is ", h.name, ", call me on ", h.phone)
}

func (h Human) sing(lyrics string) {
	fmt.Println("lalala~~~", lyrics)
}

func (h Human) guzzle(bearStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", bearStein)
}

func (e Employee) sayHi() {
	fmt.Println("Hello, my name is ", e.name, ".")
	fmt.Println("I'm working at ", e.company, ".")
	fmt.Println("Call me on ", e.phone, ".")
}

func (e Employee) spendSalary(amount float32) {
	e.money -= amount
}

func (k Kodomo) borrowMoney(amount float32) {
	k.loan += amount
}

//定义interface
type Men interface {
	sayHi()
	sing(lyrics string)
	guzzle(bearStein string)
}

type YoungChap interface {
	sayHi()
	sing(lyrics string)
	borrowMoney(amount float32)
}

type ElderlyGent interface {
	sayHi()
	sing(lyrics string)
	spendSalary(amount float32)
}

func testInterface() {
	mike := Kodomo{Human{"Mike", 22, "123123123"}, "MIT", 0.00}
	paul := Kodomo{Human{"Paul", 24, "322352352"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 27, "73457345"}, "Go", 5550}
	tom := Employee{Human{"Tom", 37, "5559867865"}, "Things", 12555}

	//定义Men类型变量i
	var i Men

	//i能存储Kodomo
	i = mike
	i.sayHi()
	i.sing("hehehe")
	i.guzzle("vodka")

	fmt.Println("--------------------------------------------------")

	//i能存储Employee
	i = tom
	i.sayHi()
	i.sing("heiheihei")
	i.guzzle("osake")

	fmt.Println("--------------------------------------------------")

	//定义slice Men
	x := make([]Men, 3)
	x[0], x[1], x[2] = mike, paul, sam
	for _, value := range x {
		value.sayHi()
	}
}
