package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	var a, b int = 0, 1
	return func() int {
		defer func(a, b *int) {
			*a, *b = *b, *a+*b
		}(&a, &b)
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
