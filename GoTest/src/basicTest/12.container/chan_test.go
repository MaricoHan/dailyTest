package container

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	c := make(chan int, 12)
	c <- 1
	c <- 2
	fmt.Println(len(c))
	close(c)
	for a := range c {
		fmt.Println(a, len(c),&a)
	}
	ss:="2134567777"
	for _,s:=range ss{
		fmt.Println(s)
	}

}
func TestNil(t *testing.T) {
	a := make(map[string]map[string]*int)
	b := 12
	a["aaa"] = make(map[string]*int)
	a["aaa"]["bbb"] = &b
	fmt.Println(*a["aaa"]["bbb"])

	var c interface{}=a
	fmt.Printf("%v",c)


}