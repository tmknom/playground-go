package main

import (
	"fmt"
	"sort"
)

func main() {
	var nowhere *int
	fmt.Println(nowhere)
	fmt.Println(&nowhere)
	if nowhere != nil {
		fmt.Println(*nowhere)
	} else {
		fmt.Println("nowhere is nil")
	}

	fmt.Println("")
	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()

	fmt.Println("")
	var fn func(a, b int) int
	fmt.Println(fn == nil)

	fmt.Println("")
	food := []string{"onion", "carrot", "celery"}
	fmt.Println(food)
	sortStrings(food, nil)
	fmt.Println(food)
	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) })
	fmt.Println(food)

	fmt.Println("")
	var soup []string
	fmt.Println(soup == nil)
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	fmt.Println(len(soup))
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println("")
	soup2 := mirepoix(nil)
	fmt.Println(soup2)

	fmt.Println("")
	var soup3 map[string]int
	fmt.Println(soup3)
	fmt.Println(soup3 == nil)
	if mesurement, ok := soup3["onion"]; ok {
		fmt.Println(mesurement)
	}
	for ingredient, mesurement := range soup3 {
		fmt.Println(ingredient, mesurement)
	}

	fmt.Println("")
	var v interface{}
	fmt.Printf("%T %v %v \n", v, v, v == nil)

	fmt.Println("")
	var p *int
	v = p
	fmt.Printf("%T %v %v \n", v, v, v == nil)
	fmt.Printf("%#v \n", v)

	fmt.Println("")
	n := newNumber(42)
	fmt.Println(n)
	e := number{}
	fmt.Println(e)
}

func (n number) String() string {
	if !n.valid {
		return "未設定"
	}
	return fmt.Sprintf("%d", n.value)
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

type number struct {
	value int
	valid bool
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

type person struct {
	age int
}
