package main

import "fmt"

//break终止
//func main()  {
//	for i:=0;i<10;i++{
//		if i==5{
//			break
//		}
//		fmt.Println(i)
//	}
//	fmt.Println("结束")
//}

//continue跳过
func main()  {
	for i:=0;i<10;i++{
		if i==4{
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("结束")
}