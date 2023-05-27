package repository

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"zoomer/internal/models"
)

type resourceRepository struct {
}

func NewResourceRepository() ResourceRepository {
	return &resourceRepository{}
}

func (rr *resourceRepository) CreateResource(id, name string) (jsonData []byte, todo models.Resource) {
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

func ImageToByte(img string) []byte {
	file, err := os.Open(img)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer := bufio.NewReader(file)
	return StreamToByte(buffer)
}

func (rr *resourceRepository) GetResource(jsonFile io.Reader) (temp models.Resource) {
	data := StreamToByte(jsonFile)
	json.Unmarshal(data, &temp)
	return temp
}

func (rr *resourceRepository) GetResourcesList(jsonFiles []io.Reader) (temp models.ResourceList) {
	for i := 0; i < len(jsonFiles); i++ {
		data := StreamToByte(jsonFiles[i])
		var todo models.Resource
		json.Unmarshal(data, &todo)
		temp.ResourceList = append(temp.ResourceList, todo)
	}
	return temp
}
