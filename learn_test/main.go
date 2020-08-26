package main //必须有个main包
import (
	"fmt"
	"time"
)

func main() {
	var i int = 10
	go func() {
		for j := 0; j < 10; j++ {
			i += 1
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Printf("i=%d\n", i)

}
