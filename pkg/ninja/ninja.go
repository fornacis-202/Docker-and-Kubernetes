package ninja

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Ninja struct {
	APIkey string `koanf:"api_key"`
}

func (n Ninja) GetWeather(city string) (min int, max int, erro error) {

	url := "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=" + city

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", n.APIkey)
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error getting data from ninja.")
		return 0, 0, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	var response Response
	json.Unmarshal(body, &response)
	return response.MinTemp, response.MaxTemp, nil

}

type Response struct {
	MinTemp int `json:"min_temp"`
	MaxTemp int `json:"max_temp"`
}
