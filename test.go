package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	/*u1 := user.User{}
	u1.Name = "zs"
	u1.Password = "ll"
	fmt.Println(u1)
	u1.Test2()*/
	test5()
}

func testGo() {
	message := make(chan string, 3)

	go func() {
		message <- "send"
		fmt.Println("go func")
	}()
	fmt.Println("main func")
	str := <-message
	fmt.Println(str)
	test1(func() {
		fmt.Println("fp3")
	})
	fn := func() {
		fmt.Println("fn")

	}
	go fn()

	time.Sleep(time.Nanosecond * 30000)
}

func test1(f1 func()) {
	fmt.Println("t1")
	f1()
}

func test() {
	/*fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// 布尔型，还有你想要的逻辑运算符。
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!trfmt.Println(n)ue)*/
	//fmt.Println(n)fmt.Println(n)
	var a = 12
	s := 121
	fmt.Println(s)
	//fmt.Println(n)
	fmt.Println(a)
	for i := 0; i < 100; i++ {
		//n := rand.Float64()
		//fmt.Println(n)
	}
	i, _ := strconv.ParseInt("1233", 0, 64)
	fmt.Print(i)
}
func test5() {
	message := make(chan string, 2)
	message <- "one"
	message <- "two"
	close(message)
	fmt.Println(cap(message))
	for e := range message {
		fmt.Println(e)
	}
}
