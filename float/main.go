package main

import (
	"fmt"
	"strings"
)

//浮点型
func main()  {
	s1:=`第一行
第二行
第三行`
	f1 :=1.232321
	f3:=`D:\linux就该这么学\go\老男孩带你21周搞定Go语言\golang课件\day01`
	f4:="abdbasdasd"
	fmt.Printf("%T\n",f1)
	f2:=float32(1.5)
	fmt.Printf("%T\n",f2)
	fmt.Println(s1)
	//拼接
	name:="理想+"
	world:="理想"
	s2:=name+world
	fmt.Println(s2)
	fmt.Println(f3)
	fmt.Println(strings.Index(f4,"b"))
	//分割
	ret:=strings.Split(f3,"\\")
	fmt.Println(ret)
	//拼接
	fmt.Println(strings.Join(ret,"+"))
	fmt.Println(len(f4))
}