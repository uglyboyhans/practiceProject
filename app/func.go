package main

import (
	"math/rand"
	"fmt"
	"time"
)

type testInt func(int) bool //声明一种函数类型，其入参为int,返回值为bool
func isOdd(a int) bool {
	return a%2 == 0
}
func isEven(a int) bool {
	return a%2 != 0
}
func filter(slice []int, fun testInt) (result []int) {
	for _,value := range slice{
		if fun(value) {
			result = append(result, value)
		}
	}
	return result
}

//if else, for
func testFlow() {
	//if else
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	age := r.Int31n(36)
	if age > 18 {
		fmt.Printf("%v歳は大人だ\n", age)
	} else if age == 18 {
		fmt.Printf("%v歳はちょうど成年者です\n", age)
	} else {
		fmt.Printf("%v歳は未成年者である\n", age)
	}

	//for
	for i := 0; i < 9; i++ {
		fmt.Printf("%v ", i*i)
	}
	fmt.Printf("\n")

	for true { // 「while」の使い方
		source1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(source1)
		age1 := r1.Int31n(36)
		if age1 > 18 {
			fmt.Printf(" %v", age1)
			break
		} else if age1 == 18 {
			continue
		} else {
			fmt.Printf(" %v", age1)
		}
	}

	fmt.Printf("\n")
}

// defer,pointer
func deferTest(a *int) int {
	//defer是采用后进先出模式
	defer println("１番")
	defer println("２番")
	defer println("３番")
	*a = 1 + *a
	return *a
}

//two return value
func SumAndProduct(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return
}

//普通の関数
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
