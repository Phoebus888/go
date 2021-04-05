package main

import "fmt"

//数组定义
func main()  {
	//根据初始值自动判断数组长度是多少，常用
	a1:=[...]int{0,1,2,3,4,5,6,7,8}
	a2:=[...]string{"北京","上海","广州"}
	fmt.Println(a1)
	fmt.Println(a2)
}

//数组遍历
//func main()  {
//	//根据索引遍历
//	area:=[...]string{"北京","上海","广州"}  //索引0-2 area[0] area[1] area[2]
//	for i:=0;i<len(area);i++{
//		fmt.Println(area[i])
//	}
//	//for range遍历
//	for _,v:=range area{
//		fmt.Println(v)
//	}
//}

//二位数组的定义
//func main()  {
//	a:=[3][2]string{
//		{"北京","上海"},
//		{"广州","深圳"},
//		{"成都","重庆"},
//	}
//	fmt.Println(a)
////多维数组的遍历
//	for _,v1:=range a{    //_循环中代表不需要这个值
//		fmt.Println(v1)
//		for _,v2:=range v1{
//			fmt.Println(v2)
//		}
//	}
////数组是值类型
//b1:=[...]int{1,2,3}
//b2:=b1
//b2[0]=100
//fmt.Println(b1,b2)
//}

//求出数组的和
//func main()  {
//	a1:=[...]int{1,3,5,7,9}
//	//定义一个值,用于和循环得出的a1数组值相加
//	sum:=0
//	for _,v1:=range a1{
//		sum=sum+v1
//	}
//	fmt.Println(sum)
////找出和为8的下标为(0,3)和(1,2)
//  for i:=0;i<len(a1);i++{
//  	for j:=i+1;j<len(a1);j++{
//  		if a1[i]+a1[j]==8{
//  			fmt.Printf("(%v,%v)\n",i,j)
//		}
//	   }
//  }
//}

