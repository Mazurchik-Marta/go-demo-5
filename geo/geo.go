package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}
type CityPopulationResponce struct {
	Error bool `json:"error"`
}

var ErrorNoCity = errors.New("NOCITY")
var ErrorNOT200 = errors.New("NOT200")
	/* Функция GetMyLocation выполняет определение местоположения пользователя.
	Определение местоположения через IP
	Если город не указан, выполняется HTTP-запрос к API https://ipapi.co/json/, 
	*/
func GetMyLocation(city string) (*GeoData, error){
	if city !="" {
		isCity := checkCity(city) 
		if !isCity {
			return nil, ErrorNoCity
		}
		return &GeoData{
			City: city,
		}, nil
	}
	/* Отправляет HTTP GET-запрос к https://ipapi.co/json/ — 
	API, которое возвращает информацию о местоположении на основе IP-адреса.
	resp — объект *http.Response, содержащий HTTP-ответ от сервера.
 	*/
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrorNOT200
	}
	defer resp.Body.Close()
	/*Читает весь поток данных resp.Body (тело HTTP-ответа).
	resp.Body — это поток данных, который нужно полностью прочитать перед закрытием.
	После чтения данных resp.Body должен быть закрыт (defer resp.Body.Close()).
	возвращает: 
	body — байтовый массив с содержимым тела ответа.
	*/
	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil

}
	/* Функция checkCity(city string) bool выполняет проверку 
	существования указанного города.
	Возвращает true, если город существует, и false, если его нет или произошла ошибка запроса. 
	*/
func checkCity(city string) bool {
	/* Создается JSON-объект с ключом "city" и значением переданного аргумента city. 
	Затем этот объект сериализуется в postBody.	
 	*/
	postBody, _ := json.Marshal(map[string]string{
		"city" : city,
	}) 
	/*
	Выполняется HTTP-запрос с заголовком Content-Type: application/json, передавая тело запроса в виде JSON.
 	Если произошла ошибка при отправке запроса (например, недоступность API), функция сразу возвращает false.
 	*/
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities","application/json",bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	/*Декодирование JSON-ответа
	Ответ сервера парсится в структуру CityPopulationResponce.
	Если поле Error в ответе false, значит город найден, и функция возвращает true. 
	Если Error == true, значит город отсутствует или произошла ошибка, и функция возвращает false.
	*/
	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error 
}
