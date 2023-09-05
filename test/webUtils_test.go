package infoplus_tests

import (
	"infoplus/server/utils"
	"reflect"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	jsonTests := []struct {
		url  string
		want map[string]interface{}
	}{
		{"https://jsonplaceholder.typicode.com/todos/1", map[string]interface{}{"userId": 1.0, "id": 1.0, "title": "delectus aut autem",
			"completed": false}},
		{"https://dummyjson.com/products/1", map[string]interface{}{
			"id":                 1.0,
			"title":              "iPhone 9",
			"description":        "An apple mobile which is nothing like apple",
			"price":              549.0,
			"discountPercentage": 12.96,
			"rating":             4.69,
			"stock":              94.0,
			"brand":              "Apple",
			"category":           "smartphones",
			"thumbnail":          "https://i.dummyjson.com/data/products/1/thumbnail.jpg",
			"images": []interface{}{
				"https://i.dummyjson.com/data/products/1/1.jpg",
				"https://i.dummyjson.com/data/products/1/2.jpg",
				"https://i.dummyjson.com/data/products/1/3.jpg",
				"https://i.dummyjson.com/data/products/1/4.jpg",
				"https://i.dummyjson.com/data/products/1/thumbnail.jpg",
			},
		}},
	}

	for _, tt := range jsonTests {
		got, err := utils.JsonDecode(tt.url)
		if err != nil {
			t.Errorf("Failed to decode, recieved %s", err)
		}
		eq := reflect.DeepEqual(got, tt.want)
		if !eq {
			t.Errorf("got %s want %s", got, tt.want)
		}
	}
}
