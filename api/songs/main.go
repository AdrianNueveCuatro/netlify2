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
		"title": "Song 1",
		"edition": "9th",
		"copyright": 2012,
		"language": "ENGLISH",
		"pages": 976,
		"author": "Abraham Silberschatz",
		"author_id": 1,
		"publisher": "John Wiley & Sons",
		"publisher_id": 1
	},
	{
		"_id": 2,
		"title": "Song 2",
		"edition": "6th",
		"copyright": 2010,
		"language": "ENGLISH",
		"pages": 1376,
		"author": "Abraham Silberschatz",
		"author_id": 1,
		"publisher": "John Wiley & Sons",
		"publisher_id": 1
	},
	{
		"_id": 3,
		"title": "Song 3",
		"edition": "5th",
		"copyright": 2010,
		"language": "ENGLISH",
		"pages": 960,
		"author": "Andrew S. Tanenbaum",
		"author_id": 2,
		"publisher": "Pearson Education",
		"publisher_id": 2
	},
	{
		"_id": 4,
		"title": "Song 4",
		"edition": "4th",
		"copyright": 2014,
		"language": "ENGLISH",
		"pages": 1136,
		"author": "Andrew S. Tanenbaum",
		"author_id": 2,
		"publisher": "Pearson Education",
		"publisher_id": 2
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
