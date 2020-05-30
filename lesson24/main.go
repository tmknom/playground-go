package main

import (
	"fmt"
	"strings"
	"time"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

func stardate(t time.Time) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type stardater interface {
	YearDay() int
	Hour() int
}

func stardate2(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v : %v", l.lat, l.long)
}

func main() {
	var t talker

	t = martian{}
	fmt.Println(t.talk())
	shout(t)

	fmt.Println("")
	t = laser(3)
	fmt.Println(t.talk())
	shout(t)

	fmt.Println("")
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)

	fmt.Println("")
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Println(day)
	fmt.Println(stardate(day))
	fmt.Println(stardate2(day))

	fmt.Println("")
	sol := sol(1422)
	fmt.Println(stardate2(sol))

	fmt.Println("")
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
}
