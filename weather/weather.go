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
/*
Функция GetWeather предназначена для получения прогноза погоды с сервиса wttr.in на основе переданных географических данных и формата ответа. Разберем ее работу поэтапно.
*/
func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}
	// Формирование базового URL
	/*
	Здесь создается URL для запроса к API wttr.in.
"https://wttr.in/" + geo.Сity формирует строку URL, добавляя название города.
url.Parse разбирает строку в объект *url.URL.
Если в процессе разбора произошла ошибка (например, передано некорректное название города), функция возвращает "ERROR_URL".
 */
	baseUrl, err := url.Parse("https://wttr.in/" + geo.Сity)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}
//  Добавление параметров запроса
	/*
Создается объект params типа url.Values, который хранит параметры запроса.
params.Add("format", fmt.Sprint(format)) добавляет параметр format, приводя format к строке (fmt.Sprint(format)).
baseUrl.RawQuery = params.Encode() кодирует параметры в URL (например, ?format=2), добавляя их к baseUrl.	
 */
	params := url.Values{}
	params.Add("format", fmt.Sprint(format)) 
	baseUrl.RawQuery = params.Encode()
	/*
Отправка HTTP-запроса
http.Get отправляет GET-запрос по сформированному URL.
Если запрос не удался (например, нет интернета или сервер недоступен), функция возвращает ошибку "ERROR_HTTP".
 */
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}
	// Чтение тела ответа
	/*
io.ReadAll(resp.Body) читает весь ответ сервера в body.
Если чтение не удалось, возвращается ошибка "ERROR_ReadBody".	
Преобразует содержимое body в строку и возвращает его.
Если ошибок нет, в body будет строка с прогнозом погоды в заданном формате.
 */
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_ReadBody")
	}
	return string(body), nil

}
