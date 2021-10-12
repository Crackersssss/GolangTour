package main

import "fmt"

func main() {

	i, j := 42, 2701

	p := &i //指向 i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21 //通过指针设置i的值
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
