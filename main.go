package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var name string
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 40

	celsium := rand.Intn(max-min+1) + min

	fmt.Print("Enter your city: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Weather in %s is %d degrees", name, celsium)
}
