package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type Song struct {
	Id           int    `json:"_id"`
	Title        string `json:"title"`
	Edition      string `json:"edition"`
	Copyright    int    `json:"copyright"`
	Language     string `json:"language"`
	Pages        int    `json:"pages"`
	Author       string `json:"author"`
	Author_Id    int    `json:"author_id"`
	Publisher    string `json:"publisher"`
	Publisher_Id int    `json:"publisher_id"`
}

var songs []Song

var jsonData string = `[
	{
		"_id": 1,
		"title": "Jennifer",
		"edition": "1",
		"copyright": 1999,
		"language": "Inglés",
		"pages": 976,
		"author": "Eurythmics",
		"author_id": 1,
		"album": "Sweet Dreams Album",
		"album_id": 1
	},
	{
		"_id": 2,
		"title": "Sweet Dreams",
		"edition": "6th",
		"copyright": 1999,
		"language": "Ingles",
		"pages": 1376,
		"author": "Eurythmics",
		"author_id": 1,
		"album": "Sweet Dreams Album",
		"album_id": 1
	},
	{
		"_id": 3,
		"title": "Blue",
		"edition": "5th",
		"copyright": 2000,
		"language": "Ingles",
		"pages": 960,
		"author": "Eiffel 65",
		"author_id": 2,
		"album": "Europop",
		"album_id": 2
	},
	{
		"_id": 4,
		"title": "Dub In Life",
		"edition": "4th",
		"copyright": 2000,
		"language": "Ingles",
		"pages": 1136,
		"author": "Eiffel 65",
		"author_id": 2,
		"album": "Europop",
		"album_id": 2
	}
]`

func FindSong(id int) *Song {
	for _, song := range songs {
		if song.Id == id {
			return &song
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(songs)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			song := FindSong(param)
			if song != nil {
				data, _ = json.Marshal(*song)
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
	_ = json.Unmarshal([]byte(jsonData), &songs)
	lambda.Start(handler)
}
