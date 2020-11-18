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
		"format": "MP3",
		"copyright": 1983,
		"language": "Inglés",
		"duration": 3,
		"author": "Eurythmics",
		"author_id": 1,
		"album": "Sweet Dreams Album",
		"album_id": 1
	},
	{
		"_id": 2,
		"title": "Sweet Dreams",
		"format": "MP3",
		"copyright": 1983,
		"language": "Inglés",
		"duration": 3,
		"author": "Eurythmics",
		"author_id": 1,
		"album": "Sweet Dreams Album",
		"album_id": 1
	},
	{
		"_id": 3,
		"title": "Blue",
		"format": "MP3",
		"copyright": 2010,
		"language": "Inglés",
		"duration": 4,
		"author": "Eiffel 65",
		"author_id": 2,
		"album": "Europop",
		"album_id": 2
	},
	{
		"_id": 4,
		"title": "Dub In Life",
		"format": "MP3",
		"copyright": 2014,
		"language": "Inglés",
		"duration": 3,
		"author": "Eiffel 65",
		"author_id": 2,
		"album": "Europop",
		"album_id": 2
	},
	{
		"_id": 5,
		"title": "Scare me",
		"format": "MP3",
		"copyright": 2013,
		"language": "Inglés",
		"duration": 3,
		"author": "Major Lazer",
		"author_id": 3,
		"album": "Free the Universe",
		"album_id": 3
	},
	{
		"_id": 6,
		"title": "Get free",
		"format": "MP3",
		"copyright": 2013,
		"language": "Inglés",
		"duration": 3,
		"author": "Major Lazer",
		"author_id": 3,
		"album": "Free the Universe",
		"album_id": 3
	},
	{
		"_id": 7,
		"title": "Rakata",
		"format": "MP3",
		"copyright": 2005,
		"language": "Español",
		"duration": 3,
		"author": "Wisin & Yandel",
		"author_id": 4,
		"album": "Pal Mundo",
		"album_id": 4
	},
	{
		"_id": 8,
		"title": "Sin el",
		"format": "MP3",
		"copyright": 2005,
		"language": "Español",
		"duration": 3,
		"author": "Wisin & Yandel",
		"author_id": 4,
		"album": "Pal Mundo",
		"album_id": 4
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
