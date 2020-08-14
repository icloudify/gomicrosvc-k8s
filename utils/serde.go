package utils

import (
	"encoding/json"
	"github.com/ravindra031/gomicrosvc-k8s/model"
)

/*type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}*/

func ToJson(b model.Book) []byte {
	toJson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return toJson
}

func FromJSON(str string) *model.Book {
	book := model.Book{}
	ba := []byte(str)
	err := json.Unmarshal(ba, &book)

	if err!=nil {
		panic(err)
	}

	return &book
}
