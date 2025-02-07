package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

// Если все хорошо то нужно вернуть строку.
func TestGetWeather(t *testing.T) {
	expected := "Longon"
	geoData := geo.GeoData{
		Сity: expected,
	}
	format := 3 // В 3 формате выдает город.
	results, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	if !strings.Contains(results, expected) {
		t.Errorf("Ожидалось %v, получение %v", expected, results)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -3},
}

// Проверка неправильный формат. Негативный тест.
// Тест неверной передачи  + Группа теста
// Группа тестов
func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // контекст
			expected := "Longon"
			geoData := geo.GeoData{
				Сity: expected,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v, получение %v", weather.ErrWrongFormat, err)
			}
		})
	}

}
