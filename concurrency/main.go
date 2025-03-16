package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const apiKey = "ecfd8aa66dac451426df9b4f6a8ef94e"

func fetchingWeather(city string, ch chan<- string, wg *sync.WaitGroup) any {
	var data struct {
		/*
		* "main": { "temp": 290.2 }
		 */
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}
	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error fetching weather for %s: %s\n", city, err)
		return data
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("error decoding weather data for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("this is the %s: %+v", city, data.Main)
	return data
}

func main() {
	startTime := time.Now()

	cities := []string{"Tehran", "London", "Toronto", "Paris", "Tokyo"}

	// ch := make(chan string, 4)
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchingWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("this operation took:", time.Since(startTime))
}
