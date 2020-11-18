package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type SongRef struct {
	SongId int    `json:"song_id"`
	Title  string `json:"title"`
}

type Author struct {
	Id          int       `json:"_id"`
	Author      string    `json:"author"`
	Nationality string    `json:"nationality"`
	BirthYear   int       `json:"birth_year"`
	Fields      string    `json:"fields"`
	Songs       []SongRef `json:"songs"`
}

var items []Author

var jsonData string = `[
	{
		"_id": 1,
		"author": "Eurythmics",
		"nationality": "Inglaterra",
		"debut": 1980,
		"fields": "Jennifer, Sweet Dreams",
		"songs": [
			{
				"song_id": 1,
				"title": "Jennifer"
			},
			{
				"song_id": 2,
				"title": "Sweet Dreams"
			}
		]
	},
	{
		"_id": 2,
		"author": "Eiffel 65",
		"nationality": "Italia",
		"debut": 1998,
		"fields": "Blue, Dub In Life",
		"songs": [
			{
				"song_id": 3,
				"title": "Blue"
			},
			{
				"song_id": 4,
				"title": "Dub In Life"
			}
		]
	},
	{
		"_id": 3,
		"author": "Major Lazer",
		"nationality": "Jamaica",
		"debut": 2008,
		"fields": "Scare me, Get free",
		"songs": [
			{
				"song_id": 5,
				"title": "Scare me"
			},
			{
				"song_id": 6,
				"title": "Get free"
			}
		]
	},
	{
		"_id": 4,
		"author": "Wisin & Yandel",
		"nationality": "Puerto Rico",
		"debut": 1998,
		"fields": "Rakata, Sin el",
		"songs": [
			{
				"song_id": 7,
				"title": "Rakata"
			},
			{
				"song_id": 8,
				"title": "Sin el"
			}
		]
	}
]`

func FindItem(id int) *Author {
	for _, item := range items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(items)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			item := FindItem(param)
			if item != nil {
				data, _ = json.Marshal(*item)
			} else {
				data = []byte("error\n")
			}
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	_ = json.Unmarshal([]byte(jsonData), &items)
	lambda.Start(handler)
}
