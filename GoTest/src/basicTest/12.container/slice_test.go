package container

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := a[2:4:6] // data[low,high,max]
	// 由于b是基于a截取出来的，所以和a在底层共用同一个数组
	b[0] = 33
	fmt.Println(b, len(b), cap(b)) // 对于截取操作，子切片个父切片共用底层数组，且cap默认到底层数组的末尾
	fmt.Println(a)                 // 当改变b的值，则a对应位置上的值也会改变

	s1 := a[2:5]
	s2 := s1[2:6:7]
	fmt.Println(s1, len(s1), cap(s1)) //cap=10-2
	fmt.Println(s2, len(s2), cap(s2)) //cap=(2+max)<10?(2+max):10-(2+2)=9-4=5

}
func TestAppend(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 := s[2:5]
	s2 := s1[2:6:7]
	fmt.Println("s:", s, len(s), cap(s))
	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2", s2, len(s2), cap(s2))

	fmt.Println("append是直接对该切片的底层数组操作，在其len()+1的位置直接赋值")
	// 当容量充足时，会直接对原来的底层数组操作，并不会重新开辟一块内存
	// 而只有当容量不够时，才会重新开辟内存

	tmp := append(s1, 11111)
	fmt.Println("tmp", tmp, len(tmp), cap(tmp))
	// 如果append接收方不是自己，则会脱离原底层数组，即s1脱离了原来的底层数组
	// 因为append是直接对底层数组操作，而append的返回值，是已经操作后的结果，即底层数组已经被修改。
	// 所以接收值tmp就是直接指向该底层数组，与s和s2共用同一个底层数组
	// 而虽然s1本来和s和s2共用一个底层数组，但是append默认不改变源切片（传入的切片），即认为使用者并不想修改s1。
	// 所以为了保证源切片不受影响，就重新开辟一块内存
	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2", s2, len(s2), cap(s2))
	fmt.Println("s", s, len(s), cap(s)) // 所有共用同一个底层数组的切片都会受影响

	fmt.Println("---------", "append没有超过容量")
	// s1 len 3 cap 8 ,因为该切片在底层数组的剩余容量还有5，所以在s1后面append五个及以内的元素，会保持在原数组上操作
	s1 = append(s1, 1111, 222, 333, 444, 555) //
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s, len(s), cap(s)) // 所有共用同一个底层数组的切片都会受影响

	fmt.Println("---------------", "当append超过容量的元素时,会离开原来的底层数组，自己重新开辟一块内存")
	s1 = append(s1, 777)
	fmt.Println(s1, len(s1), cap(s1)) // 容量扩展为原来的2倍，即2*8
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s, len(s), cap(s))

	runtime.NumCPU()
}
