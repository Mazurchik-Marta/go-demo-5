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
type CityPopulationResponce struct {
	Error bool `json:"error"`
}

// Отдельные ошибки
var ErrorNoCity = errors.New("NOCITY")
var ErrorNOT200 = errors.New("NOT200")

func GetMyLocation(city string) (*GeoData, error){
	if city !="" {
		isCity := checkCity(city) 
		if !isCity {
			return nil, ErrorNoCity
		}
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
		return nil, ErrorNOT200
	}
	defer resp.Body.Close()
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
// Сервис 2
// Проверка локации
func checkCity(city string) bool {
	/*
	func http.Post(url string, contentType string, body io.Reader) 
	(resp *http.Response, err error)
	*/
	
	// func json.Marshal(v any) ([]byte, error)
	postBody, _ := json.Marshal(map[string]string{
		"city" : city,
	}) // необходимо сформировать байты
	// func bytes.NewBuffer(buf []byte) *bytes.Buffer
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities","application/json",bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error // 
}
