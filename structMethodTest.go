package main

import (
	"fmt"
)

type Test struct {
	id int
}

type TestArray struct {
	ids []int
	mp  map[string]int
}

func testArrayTest(ta TestArray) {
	fmt.Printf("testArrayTest testArray addr :%p   ids addr %p	mp addr %p \r\n", &ta, &(ta.ids), &(ta.mp))
	ta.ids = append(ta.ids, 99)
	for i := 0; i < 20; i++ {
		ta.mp[fmt.Sprintf("id%d", i)] = i
	}
}

func testArrayTest2(ta *TestArray) {
	fmt.Printf("testArrayTest testArray addr :%p   ids addr %p	mp addr %p \r\n", ta, &(ta.ids), &ta.mp)
	for i := 0; i < 20; i++ {
		ta.mp[fmt.Sprintf("id00%d", i)] = i
	}
}

// fncW  function 回调
type fnW func(t Test) int

// funW 回调的具体实现
func (f fnW) fnWar(t Test) int {
	return f(t)
}

func (t1 *Test) showTestAdd(t *Test) int {
	i := t1.id + t.id
	fmt.Println(i)
	return i
}

func (t1 Test) showTestAdd2(t Test) int {
	i := t1.id + t.id
	fmt.Println(i)
	return i
}
func (t1 *Test) showTestAdd3(t Test) int {
	i := t1.id + t.id
	fmt.Println(i)
	return i
}

func main() {
	fmt.Println("展示值变量改变影响,值传递方式是复制值和指针别名，当修改复制的指针别名的值时，不影响老的指针的值")
	fmt.Println("array 别名append 有问题，map put get 貌似没问题")
	testArray := TestArray{ids: []int{1, 2, 3}, mp: map[string]int{"id": 1}}
	fmt.Printf("testArray addr :%p   ids addr %p mp addr %p \r\n", &testArray, &testArray.ids, &testArray.mp)
	testArrayTest(testArray)
	fmt.Println(testArray)

	testArrayTest2(&testArray)
	fmt.Println(testArray)

	fmt.Println("\n\n展示方法调用")
	fmt.Println("\t,变量实现方法,特定情况下可以调用指针方法，也可一点调用变量方法")
	tm := Test{id: 11}
	tp := Test{id: 22}
	//Test{id: 33}.showTestAdd3(tp) //Cannot call a pointer method on 'Test{id: 33}'
	tm.showTestAdd3(tp)
	(&tm).showTestAdd3(tp)

	fmt.Println("\t,变量指针实现方法，能调用指针方法，也可以调用变量方法")
	tm.showTestAdd2(tp)
	Test{id: 33}.showTestAdd2(tp) //这种方式可以
	(&tm).showTestAdd2(tp)

	fmt.Println("\t,wrapper function")
	wfun := fnW(tm.showTestAdd2)
	wfun.fnWar(tp)

	fmt.Println("\r\n展示方法变量和函数变量")
	test := &Test{id: 1}
	t2 := &Test{
		id: 2,
	}
	test.showTestAdd(t2) //方法变量

	fun := (*Test).showTestAdd //函数变量
	fun(test, t2)

	test2 := Test{id: 2}
	t22 := Test{
		id: 3,
	}
	test2.showTestAdd2(t22)   //方法变量
	fun2 := Test.showTestAdd2 //函数变量
	fun2(test2, t22)
}
