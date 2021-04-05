package main

import "fmt"

//多重判断
//func main()  {
//	if i:=2;i==1{
//	fmt.Println("大拇指")
//	}else if i==2{
//		fmt.Println("食指")
//	}else if i==3{
//		fmt.Println("中指")
//	}else {
//		fmt.Println("无效数字")
//	}
//}

//switch
//func main()  {
//	switch i:=1; i {
//	case 1,3,5,7,9:
//		fmt.Println("奇数")
//	case 2,4,6,8,10:
//		fmt.Println("偶数")
//	default:
//		fmt.Println("无效数字")
//		}
//	}

//switch表达式
func main()  {
	age:=13
	switch  {
	case age<25:
		fmt.Println("好好学习吧")
	case age >25 && age <35:
		fmt.Println("好好工作吧")
	case age>60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
