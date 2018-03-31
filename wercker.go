package main


import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Second)
		fmt.Println("Go Wercker")
	}
}
