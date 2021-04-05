package main

import (
	"fmt"
)


//常量
const (
	status_ok = 200
	notfound = 404
)

//iotas是常量计数器
//_跳过某些值
const (
	a1 = iota   //0
	a2          //1
	_           //跳过值
	a4          //3
	a5          //4
)


//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)



func main()  {
	fmt.Println(status_ok)
	fmt.Println(notfound)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(KB)
	fmt.Println(MB)
	//查看变量的类型
	fmt.Printf("%T\n",MB)
}
