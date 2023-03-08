package resources

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	Client, err := MinioClient()
	if err != nil {
		log.Println(err)
	}
	err = CreateBucket(Client, "todolist")
	if err != nil {
		log.Println(err)
	}
	//example data
	UploadJson(Client, "todolist", "todo1.json", "1", "go to school")
	UploadJson(Client, "todolist", "todo2.json", "2", "go to canteen")
	UploadJson(Client, "todolist", "todo3.json", "3", "come back home")
	//route
	router := gin.Default()
	router.GET("/api/todos", func(c *gin.Context) {
		ab := GetAllTodoss(Client, "todolist")
		c.JSON(200, ab)
	})
	router.POST("/api/todos/:uid/:id", func(c *gin.Context) {
		var todo Todo
		err := c.BindJSON(&todo)
		if err != nil {
			log.Fatal(err)
		}
		res := AddTodo(Client, "todolist", c.Param("id")+".json", c.Param("id"), todo.Name)
		c.JSON(200, res)
	})
	router.Run(":8080")
}
