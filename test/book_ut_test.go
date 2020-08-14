package test

import (
	"fmt"
	"github.com/ravindra031/gomicrosvc-k8s/model"
	"github.com/ravindra031/gomicrosvc-k8s/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_book_to_json(t *testing.T)  {
	book := model.Book{"Discovery of india", "J.L Nehru", "978-3-16-148410-0" }
	//var jsonObj []byte
	jsonObj := utils.ToJson(book)
	fmt.Println(string(jsonObj))
	assert.Equal(t, `{"title":"Discovery of india","author":"J.L Nehru","isbn":"978-3-16-148410-0"}`, string(jsonObj))
}

func Test_book_From_json(t *testing.T)  {
	str := "{\"Title\":\"Discovery of india\",\"Author\":\"J.L Nehru\",\"ISBN\":\"978-3-16-148410-0\"}"
	book := utils.FromJSON(str)
	fmt.Println(*book)
	assert.Equal(t,*book,model.Book{"Discovery of india", "J.L Nehru", "978-3-16-148410-0" })
}

func Test_book_From_json_negative(t *testing.T)  {
	str := "{\"title\":\"Discovery of india\",\"author\":\"J.L Nehru\",\"isbn\":\"978-3-16-148410-0\"}"
	book := utils.FromJSON(str)
	fmt.Println(*book)
	assert.Equal(t,*book,model.Book{"Discovery of india", "J.L Nehru", "978-3-16-148410-0" })
}

