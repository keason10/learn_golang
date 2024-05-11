package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	//error var i
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()

	//success var i,case 1
	fmt.Println()
	wg.Add(5)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()

	//success var i,case 2
	fmt.Println()
	wg.Add(5)
	for i := 0; i < 5; i++ {
		ivar := i
		go func() {
			defer wg.Done()
			fmt.Println(ivar)
		}()
	}
	wg.Wait()

	//success var i,case 3
	fmt.Println()
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(vari int) {
			defer wg.Done()
			fmt.Println(vari)
		}(i)
	}
	wg.Wait()

}
