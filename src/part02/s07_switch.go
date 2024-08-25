package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main7() {
	fmt.Println("switch 예제")

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	switch val := rng.Intn(100); {
	case val >= 50 && val < 100:
		fmt.Println("val =", val, "/ 50에서 100사이")
		fallthrough
	case val >= 25 && val < 50:
		fmt.Println("val =", val, "/ 25에서 50사이")
		fallthrough
	default:
		fmt.Println("val =", val, "/ 25이하")
	}
}
