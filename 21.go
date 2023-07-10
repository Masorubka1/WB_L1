package main

import "fmt"

// Интерфейс для унифицированного сервиса погоды
type WeatherService interface {
	GetWeather(city string) string
}

// Сервис погоды A
type WeatherServiceA struct{}

func (wsa *WeatherServiceA) GetTemperature(city string) string {
	temperature := 25
	return fmt.Sprintf("Temperature in %s (from Service A): %d°C", city, temperature)
}

// Сервис погоды B
type WeatherServiceB struct{}

func (wsb *WeatherServiceB) GetWeatherData(city string) string {
	weatherData := "Sunny"
	return fmt.Sprintf("Weather in %s (from Service B): %s", city, weatherData)
}

// Адаптер для сервиса погоды B, реализующий интерфейс WeatherService
type WeatherServiceBAdapter struct {
	weatherServiceB *WeatherServiceB
}

func (adapter *WeatherServiceBAdapter) GetWeather(city string) string {
	return adapter.weatherServiceB.GetWeatherData(city)
}

func main21() {
	// Создаем экземпляры сервисов погоды
	serviceA := &WeatherServiceA{}
	serviceB := &WeatherServiceB{}

	// Создаем адаптер для сервиса погоды B
	adapter := &WeatherServiceBAdapter{serviceB}

	// Получаем погоду с помощью сервиса A
	weatherA := serviceA.GetTemperature("London")
	fmt.Println(weatherA)

	// Получаем погоду с помощью адаптера сервиса B
	weatherB := adapter.GetWeather("Paris")
	fmt.Println(weatherB)
}
