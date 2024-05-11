package main

import (
	"fmt"
	"strconv"
)

// 函数不用返回值
func tet1(i int) (a int, er error) {
	if i == 1 {
		return
	}
	return 1, nil
}

// ... string  = []string
func loop(i int, listStr ...string) (a int, er error) {
	for _, item := range listStr {
		fmt.Println(item)
	}
	return
}
func loop1(i int, listStr []string) (a int, er error) {
	for _, item := range listStr {
		fmt.Println(item)
	}
	return
}

func callLoop1(listStr ...string) {
	loop(1, listStr[:]...)
	loop1(1, listStr[:])
}

func callLoop2(ii int, listStr ...string) {
	loop(1, listStr[:]...)
	loop1(1, listStr[:])
}

const (
	a1 int = 1 << iota
	a2
	a3
	a4
	a5
)

func main() {
	callLoop1("11", "22")
	strings := []string{"11", "22"}
	callLoop1([]string{"11", "22"}...)
	callLoop1(strings...)

	fmt.Println()
	ff := callLoop2
	ff(1, []string{"11"}...)

	fmt.Println()
	var f func(ii int) int
	if f == nil {
		fmt.Println("f is nil")
	}

	i := 28
	strconv.Itoa(i)
	strconv.Atoi("i")

	fmt.Println(fmt.Sprintf("i = %b  %o %d  %x", i, i, i, i))

	fmt.Println(fmt.Sprintf("i = %s %s %s %s", strconv.FormatInt(int64(i), 2), strconv.FormatInt(int64(i), 8),
		strconv.FormatInt(int64(i), 10), strconv.FormatInt(int64(i), 16)))

	fmt.Println(strconv.ParseInt("128", 10, 3))

	fmt.Println(a1, a2, a3, a4, a5)

	p := new(int)
	fmt.Println(p)
	*p = 1
	fmt.Println(*p)
}
