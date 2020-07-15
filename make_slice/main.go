package main

import "fmt"

func main()  {
	// make()函数创造切片
	//s3 := []int{1,3,5,7,9}
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1, len(s1), cap(s1))

	////遍历
	//for i,v := range s3 {
	//	fmt.Println(i,v)
	//}
	//append追加元素,原来的底层数组放不下的时候,Go底层就会把底层数组换一个
	s1 = append(s1, 11)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1, len(s1), cap(s1))
	ss := []int{11,22,33,44,55}
	//...表示把ss拆开
	s1 = append(s1, ss...)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1, len(s1), cap(s1))


}
