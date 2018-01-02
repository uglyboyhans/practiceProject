package main

import (
	"fmt"
	"errors"
)

//一般用var方式来定义全局变量
var s1 = "Hello Go" // type declare "string" can be omitted
func testDataType() {
	//定义变量
	var f1, f2, f3 float64
	f1 = 3.22
	f2 = 2.33
	f3 = 4.44
	sum := f1*f2 - f3
	var int3, int4 int64 = 7, 9
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v,%v\n", int3+int4, sum)
	var int1, int2 = 3, 7
	//b1, b2 := true, false
	//const PI float32 = 3.141592653
	//const MYNAME = "HaoHan"
	testSqrt(float64(int1))
	testSqrt(float64(int2))

	//定长度的数组array
	arr1 := [4]int{}
	arr1[3] = 3
	arr2 := [2][4]int{{}, {}}
	arr2[1][3] = 4

	bytes1 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'} // byte的实际类型是 uint8
	//不定长度的数组slice:
	b1 := bytes1[3:7] // 3d 4e 5f 6g不包含bytes1[7]h
	b2 := bytes1[:7]  // [0,7)
	b3 := bytes1[1:]  // [1,len()]
	//b2 = append(b2, 'z')：由于b2[7]指向bytes[7]，这句将会改变bytes1[7]的值
	//b3 = append(b3, 'z') //此时将动态分配新的数组空间
	//b4 := bytes1[:2:2]   // 容量 = 2-0，无法访问bytes1[2]以及以后的元素
	//b4 = append(b4, 'y')
	//append函数会改变slice所引用的数组的内容，
	// 从而影响到引用同一数组的其它slice。
	// 但当slice中没有剩余空间（即(cap-len) == 0）时，
	// 此时将动态分配新的数组空间。
	// 返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；
	// 其它引用此数组的slice则不受影响。
	//fmt.Printf("b4[2]:" + string(b4[2]) + "\n")
	fmt.Printf("bytes1[2]:" + string(bytes1[2]) + "\n")
	fmt.Printf("b1b2b3:%v%v%v\n", string(b1[2]), string(b2[6]), string(b3[5]))
	//slice是引用类型，试试改变bytes1里的值：
	bytes1[2] = 'i'
	bytes1[3] = 'j'
	fmt.Printf(string(b3[1]) + string(b3[2]) + "\n")

	s2 := `!!!
lalala~~~~`
	s3 := s1 + s2
	fmt.Printf(s3 + "\n")

	//复数
	c1 := 5 + 4i
	c2 := 6 + 7i
	fmt.Printf("%v\n", c1*c2)

	//error type
	err := errors.New("ダメ～something wrong")
	if err != nil {
		fmt.Println(err)
	}

	//map:
	var numbers map[string]int
	//make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
	numbers = make(map[string]int) // 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	num2 := make(map[string]int)
	num2["guigui"] = 4
	numbers["hh"] = 25
	numbers["tt"] = 24
	num2["cucu"] = 4
	num1 := num2
	fmt.Printf("cucu is %v years old\n", num2["cucu"])
	fmt.Printf("guigui is %v years old\n", num2["guigui"])
	fmt.Printf("hh is %v years old\n", numbers["hh"])
	fmt.Printf("tt is %v years old\n", numbers["tt"])
	num1["snake"] = 5 //map也是引用,此时num2["snake"]也是5
	fmt.Printf("snake is %v years old\n", num2["snake"])

}

func init(){
	fmt.Printf("This is func init\n")
}