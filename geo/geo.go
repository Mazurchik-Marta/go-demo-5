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


var ErrorNoCity = errors.New("NOCITY")
var ErrorNOT200 = errors.New("NOT200")
/*
Функция GetMyLocation выполняет определение местоположения пользователя.
Проверка переданного города (city)
Если город указан, проверяется его существование с помощью checkCity(city).
Если город не найден, возвращается ошибка ErrorNoCity.
Если город существует, функция возвращает GeoData{Сity: city}.
Определение местоположения через IP
Если город не указан, выполняется HTTP-запрос к API https://ipapi.co/json/, который возвращает данные о местоположении пользователя по его IP.
Проверяется статус-код ответа (resp.StatusCode). Если он не равен 200, возвращается ошибка ErrorNOT200.
Тело ответа читается с помощью io.ReadAll(resp.Body).
JSON-ответ десериализуется в структуру GeoData с помощью json.Unmarshal.
Возвращается указатель на GeoData и nil в качестве ошибки.

*/
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
/*
resp, err := http.Get("https://ipapi.co/json/")
Отправляет HTTP GET-запрос к https://ipapi.co/json/ — API, которое возвращает информацию о местоположении на основе IP-адреса.
Возвращает
resp — объект *http.Response, содержащий HTTP-ответ от сервера.
err — ошибку, если запрос не удался (например, нет соединения с сервером).
 */
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrorNOT200
	}
	defer resp.Body.Close()
/*
body, err := io.ReadAll(resp.Body)
Читает весь поток данных resp.Body (тело HTTP-ответа) в переменную body типа []byte.
resp.Body — это поток данных, который нужно полностью прочитать перед закрытием.
После чтения данных resp.Body должен быть закрыт (defer resp.Body.Close()).
возвращает: 
body — байтовый массив с содержимым тела ответа.
err — ошибку, если чтение не удалось.
*/
	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		return nil, err
	}
	var geo GeoData
/*
json.Unmarshal(body, &geo)	
Декодирует JSON-данные из body и записывает их в структуру geo (типа GeoData).
Использует Go-пакет encoding/json, который парсит JSON-объект и сопоставляет поля JSON с полями структуры GeoData.
Возвращает: 
Ошибку (err), если JSON-структура некорректна или несовместима с GeoData.
Заполняет структуру geo значениями из JSON.
 */
	json.Unmarshal(body, &geo)
	return &geo, nil

}
/*
Функция checkCity(city string) bool выполняет проверку существования указанного города, отправляя HTTP-запрос к API https://countriesnow.space/api/v0.1/countries/population/cities.
Возвращает true, если город существует, и false, если его нет или произошла ошибка запроса. 
*/
func checkCity(city string) bool {
	/*
Создается JSON-объект с ключом "city" и значением переданного аргумента city. Затем этот объект сериализуется в postBody.	
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
	// Чтение тела ответа:
	/*
	Закрывается resp.Body (отложенный вызов defer), затем читается ответ в body. В случае ошибки возвращается false.	
 */
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	//Декодирование JSON-ответа
	/*
Ответ сервера парсится в структуру CityPopulationResponce.
Если поле Error в ответе false, значит город найден, и функция возвращает true. Если Error == true, значит город отсутствует или произошла ошибка, и функция возвращает false.
!!

 */
	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error // 
}
