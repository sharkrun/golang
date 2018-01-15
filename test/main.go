package main

import (
	"fmt"
	"flag"
	"reflect"
	"test/util"
	"test/datatype"
)

func main() {
	etcd := flag.String("etcd", "http://127.0.0.1:2379", "etcd endpoint")
	flag.Parse()
	fmt.Println("etcd: ", *etcd)

	var a = [3]int{1,2,3}
	var b = a
	b[1]++
	fmt.Println(a, b)

	var a1 = [3]int{1,2,3}
	var b1 = &a1
	b1[1]++
	fmt.Println(a1, b1)
	fmt.Printf("b1.type: %v\n", reflect.TypeOf(b1))

	myrect0 := &util.Rect{0, 0, 100, 200}
	fmt.Println(myrect0.Area())

	myrect := util.NewRect(0, 0, 100, 200)
	fmt.Println(myrect.Area())

	datatype.Test()
}

