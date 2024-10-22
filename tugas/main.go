package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sensor(nama string, ch chan<- string) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			randomValue := rand.Intn(100)
			nilaiSensor := fmt.Sprintf("%s: %d", nama, randomValue)
			data := nilaiSensor
			ch <- data
		case <-timeout:
			fmt.Println("Sensor Timeout")
			close(ch)
			return
		}
	}
}

func main() {
	ch := make(chan string, 3)

	go sensor("Suhu", ch)
	go sensor("Kelembaban", ch)
	go sensor("Tekanan", ch)

	fmt.Println("Sensor dijalankan")

	for data := range ch {
		fmt.Println(data)
	}

}
