package main

import "fmt"

func main() {
	//多重判断
	//age:=18
	//if age> 19 {
	//	fmt.Println("已成年")
	//}else if age> 35{
	//	fmt.Println("中年")
	//}else{
	//	fmt.Println("好好学习")
	//}
	//作用域
	//age只在if判断条件中生效,减少程序内存占用
	if age := 18; age > 19 {
		fmt.Println("你已经成年")
	} else {
			fmt.Println("未成年")
	}
}