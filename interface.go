package main

import (
	"fmt"
)

//Activity ..
type Activity interface {
	Eat() string
	Run(string) bool
}

type cat struct {
	food   string
	fast   bool
	amount int
}

type dog struct {
	food string
	fast bool
}
type rat struct {
	food string
	fast bool
}

func (c cat) Eat() string {
	if c.amount < 0 {
		return c.food
	}
	ss := "emtry"
	return ss
}

func (c cat) Run(st string) bool {
	return true
}

func (d dog) Eat() string {
	return d.food
}

func (d dog) Run(string) bool {
	if d.food == "bone" {
		return true
	}
	return false
}
func (r rat) Eat() string {
	if r.food == "" {
		s := "all"
		return s
	}
	return r.food
}

func (r rat) Run(string) bool {
	return true
}
func doSomething(act Activity) {
	xType := fmt.Sprintf("%T", act)
	fmt.Println(xType)
	fmt.Printf("This food is %v\n", act.Eat())
	fmt.Printf("Running fast : %v\n", act.Run(""))
	fmt.Println()
}

func main6() {
	ca := cat{
		food:   "fish",
		amount: 9,
	}
	do := dog{food: "bone"}
	ra := rat{}
	doSomething(ca)
	doSomething(do)
	doSomething(ra)
}
