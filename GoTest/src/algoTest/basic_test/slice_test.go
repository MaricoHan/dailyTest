package basictest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(data[0 : len(data)-1])
	fmt.Println(data[1], reflect.TypeOf(&data[1]))

	tmp := data[0:4:9]
	fmt.Println(tmp,cap(tmp),len(tmp))
	tmp[5]=1
	fmt.Println(tmp)

}
func TestString(t *testing.T) {
	var s = "123"
	fmt.Println(s[:1], reflect.TypeOf(s[:1]))
	fmt.Println(string([]rune(s)[:1]))

	fmt.Println(^uint(0)-1)
}
