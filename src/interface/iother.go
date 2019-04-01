package main

import "log"

type User struct {
	Name  string
	Email string
}

// 注意这个是值接收者
func (u User) Notify() error {
	u.Name = "alimon"
	return nil
}

// 注意这个是指针接收者
func (u *User) Notify2() error {
	u.Name = "alimon"
	return nil
}

func main() {
	u := &User{"Damon", "damon@xxoo.com"}
	u.Notify()
	log.Println("u.Name", u.Name) // Damon

	u.Notify2()
	log.Println("u.Name", u.Name) // alimon
}
