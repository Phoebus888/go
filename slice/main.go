package main

import "fmt"

func main()  {
	a:=[]string{"北京","上海","南京"}
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	a1:=a[1]
	a2:=a[1:]
	a3:=a[:2]
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	//遍历切片
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
	//range遍历
	for _,v:=range a{
		fmt.Println(v)
	}
}
