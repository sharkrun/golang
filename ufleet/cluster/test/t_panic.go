package test

import "fmt"


func Tpanic() {

	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err:=recover();err!=nil{
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()

	panic("fault")
	fmt.Println("panic")
}
