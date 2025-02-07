package weather

import (
	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrWrongFormat = errors.New("ERROR_WRONG_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}
	// создать обьект url
	// func url.Parse(rawURL string) (*url.URL, error)
	// преобразует строку в полноценный url
	baseUrl, err := url.Parse("https://wttr.in/" + geo.Сity)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}
	// Добаляем в плученный url параметры.
	// type Values map[string][]string
	params := url.Values{}
	//func (v url.Values) Add(key string, value string
	params.Add("format", fmt.Sprint(format)) // Изм.1	не (string)
	//func (v url.Values) Encode() string
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_ReadBody")
	}
	return string(body), nil

}
