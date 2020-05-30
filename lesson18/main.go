package main

import (
	"fmt"
)

func main() {
	dwarfs := []string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}

	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	dwarfs = append(dwarfs, "Orcus")
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)

	terrastrial := dwarfs[0:4:4]
	dump("terrastrial", terrastrial)

	worlds := append(terrastrial, "ケレス")
	dump("worlds", worlds)
	fmt.Println(dwarfs)
	dump("dwarfs", dwarfs)

	terrastrial2 := dwarfs[0:4]
	dump("terrastrial2", terrastrial2)
	worlds2 := append(terrastrial2, "ケレス")
	dump("worlds2", worlds2)
	dump("dwarfs", dwarfs)

	dwarfs3 := make([]string, 0, 2)
	dump("dwarfs3", dwarfs3)
	dwarfs3 = append(dwarfs3, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3)

	worlds3 := terraform("New", "Venus", "Mars")
	fmt.Println(worlds3)
}

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func dump(label string, slice []string) {
	fmt.Printf("%v: 長さ %v, 容量 %v %v\n", label, len(slice), cap(slice), slice)
}
