package main

import "fmt"

func changeData(arr []int, setVal int) {
	fmt.Println(fmt.Sprintf("changeData arr addr = %p", &arr))
	for i := 0; i < len(arr); i++ {
		arr[i] = setVal
	}
}

func changeDataP(arrP *[]int, setVal int) {
	fmt.Println(fmt.Sprintf("changeDataP arr addr = %p", &arrP))
	for i := 0; i < len(*arrP); i++ {
		(*arrP)[i] = setVal
	}
}

func printAddrWithVal(arr *[]int) {
	fmt.Println(fmt.Sprintf("addr = %p,value =%v\r\n", arr, *arr))
}

func testMap(mp map[string]bool, mpP *map[string]bool) {
	fmt.Println(fmt.Sprintf("mp addr = %p , mpP addr = %p", &mp, mpP))
	mp["id"] = true
	(*mpP)["id1"] = true
}

func main() {

	//传递到函数用值传递，但是2地方所操作的基础数据是一份，函数改变数组的值，反应到main的变量上
	// map ，channel, array ,slice 都一样
	fmt.Println("复杂对象，函数值传递和指针传递实参，都是修改的同一份值")
	var intArr = []int{1, 2, 3}
	printAddrWithVal(&intArr)
	changeData(intArr, -1)
	printAddrWithVal(&intArr)
	changeDataP(&intArr, -1)
	printAddrWithVal(&intArr)

	fmt.Println("复杂对象，函数值传递和指针传递实参，map")
	//map 值传递，也和函数共享的是值所在的内存地址
	var map1 = map[string]bool{"sex": true}
	fmt.Println(fmt.Sprintf("mp addr = %p map value: %v ", &map1, map1))
	testMap(map1, &map1)
	fmt.Println(fmt.Sprintf("mp addr = %p map value: %+v ", &map1, map1))

	//复杂字段，可以初始化为nil ,简单字段不行
	fmt.Println("\r\n复杂对象nil 初始化")
	var mapVar map[string]bool = nil
	var slice []string = nil
	var ch1 chan string = nil
	var fn func() int = nil
	fmt.Println(mapVar, slice, ch1, fn)
	//var int1 int = nil // error

	fmt.Println("\r\n struct test")
	vv := struct {
		Id, Name string
	}{"id1", "name1"}
	fmt.Println(fmt.Sprintf("%T,%+v", vv, vv))
	type struct1 struct {
		Id, Name string
	}
	vv2 := struct1{"id2", "name2"}
	fmt.Println(fmt.Sprintf("%T,%+v", vv2, vv2))
}
