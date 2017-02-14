package main

import (
	"fmt"
	"time"
)

func main() {
	/*s := make([]string, 3)
	s[0] = "s"
	s[1] = "s"
	s[2] = "wq"
	s = append(s, "wos")
	for _, r := range s {
		fmt.Println(r)
	}*/
	/*c := cicle{20, 30}
	fmt.Println(add(c))*/
	/*p := person{"ls"}
	show(&p)*/
	/*for i := 0; i < 10; i++ {
		go xiancheng1()
		time.Sleep(time.Second * 1)
		go xiancheng2()
		time.Sleep(time.Second * 1)
		go xiancheng3()
		time.Sleep(time.Second * 1)
	}*/
	//tongdao()
	/*c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		c1 <- "c1"
		time.Sleep(time.Second * 1)
	}()
	go func() {
		c2 <- "c2"
		time.Sleep(time.Second * 1)
	}()
	for i := 0; i < 2; i++ {
		select {
		case ms1 := <-c1:
			fmt.Println(ms1)
		case ms2 := <-c2:
			fmt.Println(ms2)
		}
	}*/
	m := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		m <- "one"
	}()

	select {
	case a := <-m:
		fmt.Println(a)
	case <-time.After(time.Second * 2):
		fmt.Println("file")
	}

}
func niming() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type person struct {
	Name string
}

type cicle struct {
	height int
	weith  int
}

func add(c cicle) int {
	return c.height * c.weith
}
func show(p *person) {
	fmt.Println(&p)
}

type jiekou interface {
	area() int
	zhouchang() int
}

func xiancheng1() {
	fmt.Println("1开始执行了")
}
func xiancheng2() {
	fmt.Println("2开始执行了")
}
func xiancheng3() {
	fmt.Println("3开始执行了")
}
func tongdao() {
	m := make(chan string)
	go func() { m <- "ok" }()
	ms := <-m
	fmt.Println(ms)

}
