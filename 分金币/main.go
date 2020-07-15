package main

import "fmt"

/*
有5000枚金币,,分配给以下几个人 Mattehw,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth
a.名字中每包含一个e或E分1枚金币
b.名字中每包含一个i或I分2枚金币
c.名字中每包含一个o或O分3枚金币
b.名字中每包含一个u或U分4枚金币

*/
var (
	coins = 5000
	users = []string{
		"Mattehw", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	//1.一次拿到每个人的名字
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name] ++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}

	}
	//2.拿到一个人名,根据规则分金币
	//3.记录剩余金币
	left = coins

	return
}

func main()  {
	var ret int = dispatchCoin()
	fmt.Println(ret)
}
