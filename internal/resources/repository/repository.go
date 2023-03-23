package resources

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"zoomer/internal/models"
)

func CreateResource(id, name string) (jsonData []byte, todo models.Resource) {
	todo = models.Resource{
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

func GetResource(jsonFile io.Reader) (temp models.Resource) {
	data := StreamToByte(jsonFile)
	json.Unmarshal(data, &temp)
	return temp
}

func GetResourcesList(jsonFiles []io.Reader) (temp models.ResourceList) {
	for i := 0; i < len(jsonFiles); i++ {
		data := StreamToByte(jsonFiles[i])
		var todo models.Resource
		json.Unmarshal(data, &todo)
		temp.Resource = append(temp.Resource, todo)
	}
	return temp
}
