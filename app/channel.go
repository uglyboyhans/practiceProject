package main

import "fmt"

func sum(a []int, c chan<- int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func testChannel() {
	/**
    默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，
	这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，
	也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
	其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
	无缓冲channel是在多个goroutine之间同步很棒的工具。
    */
	a := []int{9, 3, 6, -2, 5, 24, -13, 4, 7}
	c := make(chan int)
	cWrite := chan<- int(c) // 类型转换
	fmt.Println("that")
	go sum(a[: len(a)/3], cWrite)
	go sum(a[len(a)/3: len(a)/3*2], cWrite)
	go sum(a[len(a)/3*2: ], cWrite)
	// 以上go同时执行完了过后（所有channel都准备好了）才会执行下一句接收
	//cRead := <-chan int(c)
	x, y, z := <-chan int(c), <-chan int(c), <-chan int(c)
	fmt.Println(x, "+", y, "+", z, "=", x+y+z)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 不close会死锁
}

func testBufferedChannel() {
	c := make(chan int, 10)
	fmt.Println(cap(c))
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciSelectVer(c, quit chan int) {
	x, y :=1,1
	for {
		select {
		case c<-x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func testSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0;i<10;i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelectVer(c, quit)
}