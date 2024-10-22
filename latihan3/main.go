package main

import (
	"fmt"
	"runtime"
	"sync"
)

//buatkan 1 func, yg menjalankan filter nilai yg ganjil untuk di tampilkan. kemudian jalankan func tadi sebanyak 10 kali di masing2 go routine. func memiliki parameter angka dan sync.waitgroup.
//

func ganjil(angka int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < angka; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println("angka ke ", i)
		}

	}
}

func main() {
	totalCore := runtime.NumCPU()
	fmt.Println("Total core :", totalCore)

	setCore := runtime.GOMAXPROCS(0)
	fmt.Println("jumlah core : ", setCore)

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		fmt.Println("------------ func ke :", i)
		ganjil(5, &wg)
	}

	wg.Wait()
	fmt.Println("program selesai")
}
