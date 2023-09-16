package main

//we import the required packages.

import (
	//"encoding/json" //for JSON decoding
	"fmt" // for output
	//"io"  // for IO functions
	//"net/http" // for HTTP requests
	//"os" // for command-line arguments
	//"github.com/faith/color" //for color to output
)

func main() {
	fmt.Println("Hello boy")
}

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
