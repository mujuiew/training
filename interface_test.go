package main

import "fmt"

type mobie interface {
	Tel(int)
	Sendsms(string) bool
}
type sumsung struct {
	num   int
	photo string
}

func (s *sumsung) Tel(i int) {
	s.num = i
}
func (s sumsung) Sendsms(st string) bool {
	return false
}

type iphone struct {
	num     int
	photo   string
	appleID int
}

func (s *iphone) Tel(i int) {
	s.num = i
}
func (s iphone) Sendsms(st string) bool {
	if s.appleID > 0 {
		return true
	}
	return false
}
func sa(mb mobie) {
	fmt.Println(mb.Sendsms("test"))
}
func main6() {
	ss := sumsung{
		num: 0,
	}
	ip := iphone{
		num:     3333,
		appleID: 9,
	}
	sa(&ss)
	sa(&ip)
	fmt.Printf("(%v, %T)\n", ip, ip)
}
