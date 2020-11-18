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
	Release   int       `json:"release"`
	Genre    string    `json:"genre"`
	Songs     []SongRef `json:"songs"`
}

var items []Album

var jsonData string = `[
	{
		"_id": 1,
		"album": "Sweet Dreams Album",
		"country": "Inglaterra",
		"release": 1983,
		"genre": "Rock",
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
		"album": "Europop",
		"country": "Italia",
		"release": 1999,
		"genre": "Pop",
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
		"album": "Free the Universe",
		"country": "Jamaica",
		"release": 2013,
		"genre": "Dancehall",
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
		"album": "Pal Mundo",
		"country": "Puerto Rico",
		"release": 2005,
		"genre": "Reggaeton",
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
