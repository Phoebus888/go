package main

import "fmt"

func main()  {
	//copy()复制切片
	a:=[]int{1,2,3,4,5}
	b:=a   //赋值
	c:=make([]int,5,6)  //定义切片，长度为5,容量为6
    copy(c,a)  //copy切片,目标切片，数据来源切片
	a[0]=100   //修改a变量索引0的值为100
	fmt.Println(a,b,c)
}
