package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	lat  float64
	long float64
}

func main() {
	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636
	fmt.Println(spirit)
	fmt.Println(spirit.lat)

	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Printf("%v\n", opportunity)
	fmt.Printf("%+v\n", opportunity)

	curiosity := opportunity
	opportunity.long += 0.0106
	fmt.Println(opportunity, curiosity)

	locations := []location{
		{lat: -1.9462, long: 354.4734},
		{lat: -14.5684, long: 175.472636},
	}
	fmt.Println(locations)

	type loc struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	curiosity2 := loc{Lat: -1.9462, Long: 354.4734}
	fmt.Println(curiosity2)
	bytes, err := json.Marshal(curiosity2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
