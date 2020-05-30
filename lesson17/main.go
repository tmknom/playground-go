package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	planets := [...]string{
		"水星",
		"金星",
		"地球",
		"火星",
		"土星",
		"木星",
		"天王星",
		"海王星",
	}

	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)

	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiantsMarkII, iceGiants)

	planetsSlice := []string{
		" 水星 ",
		"金星",
		"地球 ",
		"火星",
		"土星",
		"木星",
		"天王星",
		"海王星",
	}

	fmt.Println(planetsSlice)
	hyperspace(planetsSlice)
	fmt.Println(planetsSlice)
	hyperspace([]string{"foo", "bar"})

	sort.StringSlice(planetsSlice).Sort()
	sort.Strings(planetsSlice)
	fmt.Println(planetsSlice)
}

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
