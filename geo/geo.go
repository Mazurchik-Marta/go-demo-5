package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)


type GeoData struct {
	Сity string `json:"city"`
}

func GetMyLocotion(city string) (*GeoData, error){
	if city !="" {
		return &GeoData{
			Сity: city,
		}, nil
	}
	
	//func http.Get(url string) (resp *http.Response, err error)
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}
	//resp.Body field Body io.ReadCloser
	body, err := io.ReadAll(resp.Body) //var body []byte
	if err != nil {
		return nil, err
	}
	var geo GeoData
	// func json.Unmarshal(data []byte, v any) error
	json.Unmarshal(body, &geo)
	return &geo, nil

}
func CheckCity(city string) bool {
	/*
	func http.Post(url string, contentType string, body io.Reader) 
	(resp *http.Response, err error)
	*/
	
	// func json.Marshal(v any) ([]byte, error)
	postBody, _ := json.Marshal(map[string]string{
		"city" : city,
	}) // необходимо сформировать байты
	// func bytes.NewBuffer(buf []byte) *bytes.Buffer
	resp, err := http.Post("https://ipapi.co/json/","application/json",bytes.NewBuffer(postBody))
}


