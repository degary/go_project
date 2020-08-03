package main //必须有个main包
import (
	"fmt"
	"runtime"
)

func test(config ...string) {
	fmt.Printf("type of config is %T\n", config)
	configs := append(config, "{}")[0]
	fmt.Println(configs)
}

func main() {
	config := "abc"
	test(config)
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
}
