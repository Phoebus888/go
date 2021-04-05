package main

import "fmt"

func main()  {
	//标准循环
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	//range循环
	s:="hello-phoebus"
	for i,v:=range s{
		fmt.Printf("%d %c\n",i,v)
	}
	s1:="测试"
	fmt.Printf("%v",s1)
	//九九乘法表
//	for i:=1;i<10;i++{
//	for j:=1;j<=i;j++{
//		fmt.Printf("%d*%d=%d\t",j,i,j*i)
//	}
//	fmt.Println()
//	}
}