package user

import "fmt"

type User struct {
	Name     string
	Password string
}

func (*User) test1() {
	fmt.Println("test1")
}

func (*User) Test2() {

}
