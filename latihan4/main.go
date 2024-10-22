package main

import (
	"fmt"
	"sync"
)

// buatkan 1 struct yang memiliki properti kuantiti dan nama produk. kemudian buatkan sebuah func yang fungsinya untuk mengurangi kuantiti dari struct tsb. func ini memiliki parameter jumlah kuantiti yang akan di kurangi. jalankan func tsb di dalam go routine sebanyak 10kali. cetak kuantiti terakhirnya dari struct tsb.

type Produk struct {
	NamaProduk string
	Kuantitas  int
	mu         sync.Mutex
}

func (p *Produk) Kurang(angka int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Kuantitas -= angka
}

func main() {

	produk1 := Produk{
		NamaProduk: "Ale-ale",
		Kuantitas:  1000,
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			produk1.Kurang(10)
		}()
	}

	wg.Wait()
	fmt.Println("Kuantitas terakhir dari Ale-ale yaitu : ", produk1.Kuantitas)
}
