package main

import (
	"fmt"
	"particles/particles"
	"time"
)

func main() {
	coffee := particles.NewCoffee(5, 3)
	coffee.Start()

	timer := time.NewTicker(100 * time.Millisecond)
	for {
		<-timer.C
		fmt.Printf("\033[H\033[2J")
		coffee.Update()
		fmt.Printf("%v", coffee.Display())

	}
}
