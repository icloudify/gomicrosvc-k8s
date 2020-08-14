package main

import (
	"encoding/json"
	"fmt"
	"github.com/ravindra031/gomicrosvc-k8s/model"
	"github.com/ravindra031/gomicrosvc-k8s/utils"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("######### Service Started #########")
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", EchoFunc)
	http.HandleFunc("/api/books", BookFunc)
	http.HandleFunc("/api/books/", BookFunc)

	http.ListenAndServe(":8080", nil)

}

func BookFunc(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	switch method {
	case http.MethodGet:
		books := Allbooks()
		writer.WriteHeader(http.StatusOK)
		writer.Write(books)
	case http.MethodPost:
		ba, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		book := utils.FromJSON(string(ba))
		ab := []model.Book{}
		json.Unmarshal(Allbooks(), &ab )
		ab = append(ab, *book)
		fmt.Println(ab)
		writer.WriteHeader(http.StatusOK)
		abr,err :=json.Marshal(ab)
		if err!=nil {
			panic(err)
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(abr)

	default:
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func Allbooks() []byte {
	books := []model.Book{
		model.Book{
			Title:  "struggle of life",
			Author: "Ravindra",
			ISBN:   "123-456-234567",
		}, model.Book{
			Title:  "life in IT",
			Author: "Ravindra",
			ISBN:   "123-456-2345",
		}}

	book, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	return book

}

func EchoFunc(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Echo.....")
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Hello")
}
