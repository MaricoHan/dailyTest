package main

import (
	"fmt"
	"os"
	"sync"
)

var group sync.WaitGroup

func main() {
	//在goRoutine之前，确保被执行
	group.Add(1)

	go func() {
		defer group.Done()
		func() {
			fmt.Println("这是子go程内部的函数")
			//return //返回当前函数
			os.Exit(-1) //退出当前进程，即整个运行的程序
			//runtime.Goexit()//退出当前go程
		}()
		fmt.Println("这是子go程结束！")
	}()

	fmt.Println("这是主go程")
	group.Wait()
}
