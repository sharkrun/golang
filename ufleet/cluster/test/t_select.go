package test

import (
	"fmt"
	"time"
)

func Tselect(){
	ch := make(chan int, 1)
	for {
		select {
				case ch <- 0:
				case ch <- 1:
		}
		i := <- ch
		fmt.Println("Value received:", i)
	}
}

func Tselect1() {
	// 匿名的超时等待函数
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	fmt.Println("wait get data.....")
	select {
				case <-ch:

				case m, ok:= <-timeout:
		  			fmt.Println("get data from timeout:", m)

		  			// 关闭前的状态
		  			fmt.Println("timeout before closed status:", ok)

		  			// 	关闭channel
		  			close(timeout)

		  			// 检查channel是否关闭
		  			_, ok = <- timeout
		  			fmt.Println("timeout after closed status:", ok)
	}



}