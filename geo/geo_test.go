package geo_test // стируем только то что экспортируется!

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocotion(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "Moscow"
	expected := geo.GeoData{
		Сity: "Moscow",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата с expected
	if err != nil {
		t.Error(err)
	}
	if got.Сity != expected.Сity {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}

}
// Негативный тест
func TestTestGetMyLocotionNoCity(t *testing.T)  {
	city := "Londonsloi"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrorNoCity, err)
	}
}