package geo_test 

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocotion(t *testing.T) {
	// Arrange 
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	// Act 
	got, err := geo.GetMyLocation(city)
	// Assert 
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}

}
func TestTestGetMyLocotionNoCity(t *testing.T)  {
	city := "Londonsloi"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrorNoCity, err)
	}
}