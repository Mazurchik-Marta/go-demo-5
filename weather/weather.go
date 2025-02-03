package weather

import (
	"demo/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather (geo geo.GeoData, format int) string{
	// создать обьект url
	// func url.Parse(rawURL string) (*url.URL, error)
	// преобразует строку в полноценный url
	baseUrl, err := url.Parse("https://wttr.in/" + geo.Сity)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
// type Values map[string][]string
	params := url.Values{}
//func (v url.Values) Add(key string, value string
	params.Add("format", fmt.Sprint(format))
	//func (v url.Values) Encode() string
	baseUrl.RawQuery = params.Encode()
	resp , err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)


}