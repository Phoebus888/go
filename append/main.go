package main

import "fmt"

//append
func main()  {
	s1:=[]string{"北京","上海","广州"}
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",s1,len(s1),cap(s1))
	//append追加元素,变量必须和要添加的切片变量名一致
	//append追加元素,原来的底层数组放不下的时候,go底层会把底层数组容量*2
	s1=append(s1,"深圳")
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",s1,len(s1),cap(s1))
    s2:=[]string{"成都","武汉","苏州"}
    s1=append(s1,s2...)    //...表示拆开
	fmt.Printf("s1=%v,len(s1)=%v,cap(s1)=%v\n",s1,len(s1),cap(s1))
}
