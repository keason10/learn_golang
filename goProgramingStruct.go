package main

import (
	"fmt"
	"sync"
)

type Int1 int

// 相当于自定义了int ,扩展列int的能力
func (*Int1) printString(p int) {
	fmt.Println("printString=", p)
}

// 相当于自定义了int ,扩展列int的能力
func (Int1) printString2(p int) {
	fmt.Println("printString=", p)
}

// 相当于自定义了int ,扩展列int的能力
func (*Int1) printString3(p *int) {
	fmt.Println("printString=", p)
}

// 相当于自定义了int ,扩展列int的能力
func (Int1) printString4(p int) {
	fmt.Println("printString=", p)
}

// 变量申明的方式，2种
// golang存在4中声明，var constant ,type,func
func varDec() {
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	a1, b1, c1 := 1, 2, 3
	fmt.Println(a1, b1, c1)

	var aaa, bb, cc = 1, "2", false
	fmt.Println(aaa, bb, cc)
	a2, b2, c2 := "1", 22, true
	fmt.Println(a2, b2, c2)
}

type notifier interface {
	notify()
}

type user struct {
	X    string
	name string
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.X)
}

func sendNotification(n notifier) {
	n.notify()
}

type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.Y *= factor
	p.X *= factor
}

func main() {
	//指针，和值 都能调用方法
	p2 := &Point{1, 2}
	p2.ScaleBy(2)
	fmt.Println(*p2)
	(*p2).ScaleBy(2)
	fmt.Println(*p2)

	p3 := Point{1, 2}
	p3.ScaleBy(2)
	fmt.Println(p3)
	(&p3).ScaleBy(2)
	fmt.Println(p3)

	u := user{"email", "Bill"}
	sendNotification(u)

	var aa Int1 = 1
	aa.printString(111)
	(&aa).printString(222)
	aa.printString2(1111)
	(&aa).printString2(2222)

	var tmp = 1
	aa.printString3(&tmp)
	aa.printString4(tmp)

	var wg sync.WaitGroup
	wg.Add(2)

	//test 数据组
	fmt.Println("test 数组")
	arr := []string{"1", "2"}
	fmt.Println(fmt.Sprintf("arr addr = %p", &arr))
	go func(pa []string) {
		defer wg.Done()
		fmt.Println(fmt.Sprintf("func value %v addr1 = %p", pa, &pa))
	}(arr)
	go func(pa *[]string) {
		defer wg.Done()
		fmt.Println(fmt.Sprintf("func value %v addr2 = %p", *pa, pa))
	}(&arr)
	wg.Wait()

	fmt.Println("test slice")
	strSlice := make([]string, 2)
	wg.Add(2)
	fmt.Println(fmt.Sprintf("strSlice addr = %p", &strSlice))
	go func(pa []string) {
		defer wg.Done()
		fmt.Println(fmt.Sprintf("func value %v addr1 = %p", pa, &pa))
	}(strSlice)
	go func(pa *[]string) {
		defer wg.Done()
		fmt.Println(fmt.Sprintf("func value %v addr2 = %p", *pa, pa))
	}(&strSlice)
	wg.Wait()

}
