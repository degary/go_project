package main //必须有个main包
import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int
	Name string
}

func main() {

	users := []User{{Id: 123, Name: "user1"}, {Id: 124, Name: "user2"}}

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.Run(":8080")
}
