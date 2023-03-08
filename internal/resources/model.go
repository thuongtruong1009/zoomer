package resources

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type TodoList struct {
	TodoList []Todo `json:"TodoList"`
}
type Todo struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

func CreateJson(id, name string) (jsonData []byte, todo Todo) {
	todo = Todo{
		Id: id, Name: name,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData, todo
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
func GetATodo(jsonFile io.Reader) (temp Todo) {
	data := StreamToByte(jsonFile)
	json.Unmarshal(data, &temp)
	return temp
}
func GetAllTodos(jsonFiles []io.Reader) (temp TodoList) {
	for i := 0; i < len(jsonFiles); i++ {
		data := StreamToByte(jsonFiles[i])
		var todo Todo
		json.Unmarshal(data, &todo)
		temp.TodoList = append(temp.TodoList, todo)
	}
	return temp
}
