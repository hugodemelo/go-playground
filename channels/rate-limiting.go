package channels

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func RunRateLimiting() {
	channel := make(chan int)
	limiter := make(chan int, 3)

	go generateValue(channel, limiter)
	go generateValue(channel, limiter)
	go generateValue(channel, limiter)
	go generateValue(channel, limiter)
	go generateValue(channel, limiter)

	iterations := 0
	sum := 0
	for value := range channel {
		sum += value
		iterations++
		if iterations == 5 {
			close(channel)
		}
	}
	close(limiter)

	fmt.Println(sum)
}

func generateValue(channel chan int, limit chan int) {
	limit <- 1

	generatedNumber := randN.Intn(100)
	fmt.Printf("Go routine with %v started.\n", generatedNumber)
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("Go routine with %v finished.\n", generatedNumber)

	<-limit
	channel <- generatedNumber
}
