package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {

	fmt.Println("Новый проект!")
	city := flag.String("city", "", "Город пользователя")
	fotmat := flag.Int("fotmat", 1, "Формат вывода погоды")

	flag.Parse()
	fmt.Println(*city)
	geoData, err := geo.GetMyLocotion(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData := weather.GetWeather(*geoData, *fotmat)
	fmt.Println(weatherData)

}
