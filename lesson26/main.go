package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	answer := 42
	fmt.Println(answer)
	fmt.Println(&answer)

	fmt.Println("")
	address := &answer
	fmt.Println(address)
	fmt.Println(*address)
	fmt.Printf("address の型:%T\n", address)

	fmt.Println("")
	var address2 *int
	address2 = &answer
	fmt.Println(address2)
	fmt.Println(*address2)

	fmt.Println("")
	canada := "Canada"
	var home *string
	fmt.Printf("home の型:%T\n", home)
	home = &canada
	fmt.Println(*home)

	fmt.Println("")
	var admin *string
	scolese := "J. Scolese"
	admin = &scolese
	fmt.Println(*admin)
	bolden := "F. Bolden Jr"
	admin = &bolden
	fmt.Println(*admin)
	*admin = "Frank Bolden Jr"
	fmt.Println(*admin)

	fmt.Println("")
	major := admin
	*major = "Major"
	fmt.Println(bolden)
	fmt.Println(major == admin)
	lightfoot := "Lightfoot Jr."
	admin = &lightfoot
	fmt.Println(major == admin)

	fmt.Println("")
	fmt.Println(*major)
	charles := *major
	*major = "Bolden"
	fmt.Println(*major)
	fmt.Println(charles)
	fmt.Println(bolden)
	charles = "Bolden"
	fmt.Println(charles == bolden)
	fmt.Println(&charles == &bolden)

	fmt.Println("")
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	fmt.Println(timmy)
	timmy.superpower = "flying"
	fmt.Println(timmy)
	fmt.Println(timmy.superpower)
	fmt.Println((*timmy).superpower)

	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
	fmt.Println((*superpowers)[0])
	fmt.Println((*superpowers)[1:2])

	fmt.Println("")
	rebecca := person{
		name:       "Rebecca",
		age:        14,
		superpower: "imagination",
	}
	birthday(&rebecca)
	fmt.Println(rebecca)

	fmt.Println("")
	timmy.birthday()
	fmt.Println(timmy)

	fmt.Println("")
	nathan := person{
		name: "Nathan",
		age:  17,
	}
	nathan.birthday()
	fmt.Println(nathan)

	fmt.Println("")
	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)
	fmt.Println(day.Format(layout))
	fmt.Println(tomorrow.Format(layout))

	fmt.Println("")
	player := character{name: "Mathias"}
	fmt.Println(player)
	levelUp(&player.stats)
	fmt.Println(player)
	levelUp(&player.stats)
	fmt.Println(player)
	levelUp(&(player.stats))
	fmt.Println(player)

	fmt.Println("")
	planets := []string{
		" 水星 ",
		"金星",
		"地球 ",
		"火星",
		"土星",
		"木星",
		"天王星",
		"海王星",
	}
	fmt.Println(planets)
	reclassify(&planets)
	fmt.Println(planets)

	fmt.Println("")
	shout(martian{})
	shout(&martian{})

	fmt.Println("")
	pew := laser(2)
	shout(&pew)
}

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:4]
}

type character struct {
	name string
	stats
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type stats struct {
	level             int
	endurance, health int
}

func (p *person) birthday() {
	p.age++
}

func birthday(p *person) {
	p.age++
}

type person struct {
	name, superpower string
	age              int
}
