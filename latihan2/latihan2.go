//buatkan 3 func yg memiliki parameter channel dan msg , kemudian 3 func di jalankan di 3 go routine yg berbeda, 3 func tsb mempunyai channelnya masing2. tampilkan pesan dari 3 func yang paling cepat responnya atau mengirim pesan.

package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	pesan1 := "pesan 1"

	ch2 := make(chan string)
	pesan2 := "pesan 2"

	ch3 := make(chan string)
	pesan3 := "pesan 3"

	go msg1(ch1, pesan1)
	go msg2(ch2, pesan2)
	go msg3(ch3, pesan3)

	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received:", msg)
		case msg := <-ch2:
			fmt.Println("Received:", msg)
		case msg := <-ch3:
			fmt.Println("Received:", msg)
		}
	}

}

func msg1(ch chan string, pesan string) {
	ch <- pesan
}
func msg2(ch chan string, pesan string) {
	ch <- pesan
}
func msg3(ch chan string, pesan string) {
	ch <- pesan
}
