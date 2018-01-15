package test

import (
	"fmt"
	"time"
)

func Count(ch chan int) {

	fmt.Println("Counting", time.Now())
	time.Sleep(1e9) // sleep 1s
	ch <- 1
}


func TaskCreate(){
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		value := <- ch
		fmt.Println("value:", value)
	}
	fmt.Println("all done")
}

func Start(){
	for {
		TaskCreate()
		time.Sleep(1e9)
	}
}