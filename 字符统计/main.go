package main

import (
	"fmt"
	"strings"
	"unicode"
)
//1.判断字符串中汉字的数量
// 难点是怎么判断汉字
func main(){
	s1 := "Hello沙河"
	var count int
	for _,c := range s1 {
		if unicode.Is(unicode.Han,c) {
			count++
		}
	}
	fmt.Println(count)

	//单词出现的次数
	//把字符串按照空格切割 得到切片
	s2 := "how do you do"
	s3 := strings.Split(s2," ")
	//遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		//如果原来map中不存在w这个key,则出现次数=1
		//如果map中存在w这个key,则出现次数+1
		if _,ok := m1[w]; !ok {
			m1[w] = 1
		}else {
			m1[w]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}


