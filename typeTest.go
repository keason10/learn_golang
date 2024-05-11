package main

import (
	"fmt"
	package1 "learn_golang/p1"
	"log"
)

// 声明变量就开始分配内存地址
func showVarAddr() {
	for i := 0; i < 2; i++ {
		var a int
		var b string
		fmt.Println(&a, &b)
	}
}

func main() {
	//同一个包的变量，小写字母可以访问
	int1 := Int1(1)
	int1.printString2(11)

	test := Test{id: 1}
	test2 := TestArray{}
	fmt.Println(test.id, test2)

	p1 := package1.P1{
		//f1://在另外一个包，属于私有，这里不能显示
		F2: "11",
	}
	log.Fatalf("hello %+v %+v ", p1, &p1)
	fmt.Println(111)
	showVarAddr()
}
