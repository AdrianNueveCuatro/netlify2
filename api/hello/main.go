package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"strconv"
)

type Book struct {
	Id           int    `json:"_id"`
	Title        string `json:"title"`
	Edition      string `json:"edition"`
	Copyright    int    `json:"copyright"`
	Language     string `json:"language"`
	Pages        int    `json:"pages"`
	Author       string `json:"author"`
	Author_Id    int    `json:"author_id"`
	Publisher    string `json:"publisher"`
	Publisher_Id int    `json:publisher_id`
}

var books []Book

func FindBook(id int) *Book {
	for _, book := range books {
		if book.Id == id {
			return &book
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(req.QueryStringParameters["id"])
	var data []byte
	if err == nil {
		book := FindBook(id)
		if book != nil {
			data, _ = json.Marshal(*book)
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/html"},
		Body:            "<" + string(id) + ">",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	file, _ := ioutil.ReadFile("books.json")
	_ = json.Unmarshal([]byte(file), &books)
	lambda.Start(handler)
}
