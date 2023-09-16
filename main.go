package main

//we import the required packages.

import (
	//"encoding/json" //for JSON decoding
	"encoding/json"
	"fmt" // for output
	"io"
	"net/http" // for HTTP requests
	"os"       // for command-line arguments
	//"io"  // for IO functions
	//"github.com/faith/color" //for color to output
)

//Define weather struct
//each field will be tagged with their corresponding JSON proprety

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"current"`
	} `json:"current"`
	Forecast struct {
		Forecastady []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {

	fmt.Println("Hello boy")

	// we define the default location "Iasi"
	// This is in case no command-line arguments are provided

	q := "Iasi"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	//call to the weather API

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=fc628bc193e3475bb77152244231609&q=" + q + "&days=1&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
}
