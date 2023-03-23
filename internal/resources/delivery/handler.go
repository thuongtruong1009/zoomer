package delivery

import (
	"log"
	"time"
	"github.com/labstack/echo/v4"
)

func GetResource(bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		ab := GetAllTodoss(Client, bucketName)
		return c.JSON(200, ab)
	}
}

func CreateResource(bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		var todo Todo
		err := c.Bind(&todo)
		if err != nil {
			log.Fatal(err)
		}
		res := AddTodo(Client, bucketName, c.Param("id")+".json", c.Param("id"), todo.Name)
		return c.JSON(200, res)
	}
}


func ResourceHandler() {
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
	// UploadResource(Client, "todolist", "todo1.json", "1", "go to school")
	// UploadResource(Client, "todolist", "todo2.json", "2", "go to canteen")
	// UploadResource(Client, "todolist", "todo3.json", "3", "come back home")
}
