package main   //声明包类型

import "fmt"   //导入内置包

var user_name string //定义一个变量

var (
	name string  //字符串
	age int      //整型
	isOK bool    //布尔型
	//类型推导(根据判断变量是什么类型)
	s1="hello,world"
	s2 = 20
	s3=true
)

func main()  {
	s4:=50  //简短变量声明,只能在函数中使用
	name = "理想"
	age = 16
	isOK = true
	user_name = "phoebus"
	fmt.Printf("name:%s",name)   //%s占位符,使用name这个变量去替换占位符
	fmt.Println(age)           //打印完成之后会在后面加一个换行符
	fmt.Print(isOK)            //在终端中输出要打印的内容分
	fmt.Println(user_name)
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s4)
}