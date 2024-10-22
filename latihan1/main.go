package main

import "fmt"

// buatkan func untuk countdown , func di jalankan di go routine kemudian countdown ditampilkan di main func .

func countdown(angka int, ch chan int) {

	for i := angka; i > 0; i-- {
		ch <- i
	}
	close(ch)
}

func main() {

	chan1 := make(chan int)
	go countdown(10, chan1)

	for hasil := range chan1 {
		fmt.Println(hasil)
	}

	fmt.Println("selesai")

}
