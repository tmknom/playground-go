package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("End of", gopherID, " ...")
		case <-timeout:
			fmt.Println("我慢の限界だ！")
			return
		}
	}

	c0 := make(chan string)
	c1 := make(chan string)
	go source(c0)
	go filter(c0, c1)
	printLast(c1)

}

func printLast(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func filter(upstream, downsream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downsream <- item
		}
	}
	close(downsream)
}

func source(downsream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downsream <- v
	}
	close(downsream)
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	//fmt.Println("...", id, " snore ...")
	c <- id
}
