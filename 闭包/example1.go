package main

import (
	"fmt"
	"strings"
)
//定义个函数,判断如果不是以 ...结尾的,则返回值后加上...
func makeSuffixFunc(suffix string) func(string)string{
	return func(name string) string{
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}

func main()  {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("text"))
	fmt.Println(txtFunc("text"))

}