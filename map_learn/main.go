package main

import "fmt"

func main()  {
	//元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)

	//对内部的map进行初始化
	s1[0] = make(map[int]string,1)
	s1[0][10] = "北京"
	fmt.Println(s1)


	//值为切片类型的map
	var m1 = make(map[string][]int,10)
	m1["北京"] = []int{1,3,5,7}
	fmt.Println(m1)

}
