package main

import (
	"fmt"
)

func terraform(planets [3]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}

func main() {
	var planets [8]string
	planets[0] = "太陽"
	planets[1] = "金星"
	planets[2] = "地球"

	fmt.Println(planets[3])
	earth := planets[2]
	fmt.Println(earth)

	dwarfs := [5]string{"ケレス", "冥王星", "ハウメア"}
	fmt.Println(dwarfs[1])

	dwarfs2 := [...]string{"ケレス", "冥王星", "ハウメア"}
	fmt.Println(len(dwarfs2))

	for i := 0; i < len(dwarfs2); i++ {
		fmt.Println(i, dwarfs2[i])
	}
	for i, dwarf := range dwarfs2 {
		fmt.Println(i, dwarf)
	}
	for _, dwarf := range dwarfs2 {
		fmt.Println(dwarf)
	}

	dwarfs3 := dwarfs2
	dwarfs2[0] = "whoops"
	fmt.Println(dwarfs3)
	fmt.Println(dwarfs2)

	terraform(dwarfs3)
	fmt.Println(dwarfs3)
}
