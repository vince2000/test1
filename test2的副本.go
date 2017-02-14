package main

import (
	"errors"
	"fmt"
	"time"
)

func niming1() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	/*m := make(map[string]int)
		m["k1"] = 11
		m["k2"] = 12
		m["k3"] = 13
		m["k4"] = 14
		str := make([]string, 0)
		for key, _ := range m {
			str func niming1() func() int {
		i := 0
		return func() int {
			i += 1
			return i
		}
	}= append(str, key)
		}
		for _, r := range str {
			fmt.Println(r, m[r])
		}
		for k, v := range m {
			fmt.Println(k, v)
		}*/
	/*	i := niming1()
		fmt.Println(i())
		fmt.Println(i())
		fmt.Println(i())*/
	/*p := person{"xiaowang"}
	show2(&p)
	fmt.Println(p.Name)*/
	/*r, e := fanhui(234)
	fmt.Println(r, e)*/
	for {
		xiecheng()
	}
	//tongdao1()
}

type person struct {
	Name string
}

func show1(p person) {
	p.Name = "xiaozhang"
}
func show2(p *person) {
	p.Name = "xiaozhang"
}
func fanhui(args int) (int, error) {
	if args == 40 {
		return -1, errors.New("错误信息..")
	}
	return 1, nil
}

type jiekou interface {
	show3() int
	show4() int
}

func xiecheng() {
	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("线程1开始工作了")

	}()
	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("线程2开始工作了")

	}()
	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("线程3开始工作了")

	}()

	go func() {
		time.Sleep(time.Second * 4)
		fmt.Println("线程4开始工作了")

	}()

}
func tongdao1() {
	m := make(chan string)
	go func() {
		m <- "kaishi"
	}()
	//m <- "kaishi"
	str := <-m
	fmt.Println(str)
}
func ad2d() {
	fmt.Println("dasdasdas")
}
