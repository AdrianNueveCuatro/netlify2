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

type Album struct {
	Id        int       `json:"_id"`
	Album     string    `json:"album"`
	Country   string    `json:"country"`
	Founded   int       `json:"founded"`
	Genere    string    `json:"genere"`
	Songs     []SongRef `json:"songs"`
}

var items []Album

var jsonData string = `[
	{
		"_id": 1,
		"album": "album 1",
		"country": "United States",
		"founded": 1807,
		"genere": "Academic",
		"songs": [
			{
				"song_id": 1,
				"title": "Operating System Concepts"
			},
			{
				"song_id": 2,
				"title": "Database System Concepts"
			}
		]
	},
	{
		"_id": 2,
		"album": "album 2",
		"country": "United Kingdom",
		"founded": 1844,
		"genere": "Education",
		"songs": [
			{
				"song_id": 3,
				"title": "Computer Networks"
			},
			{
				"song_id": 4,
				"title": "Modern Operating Systems"
			}
		]
	}
]`

func FindItem(id int) *Album {
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
