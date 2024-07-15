package goroutinepzngo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) { // pool di go aman dari race condition
	pool := sync.Pool{
		New: func() interface{} { // kalo gamau ngeprint/ambil nilai nil, bisa override attr New/bikin fungsi returnnya any/interface{} biar sbg default value
			return "New"
		},
	}

	pool.Put("Luthfi")
	pool.Put("Izzuddin")
	pool.Put("Hanif")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get() // setelah diambil, data harus dikembaliin ke pool karena setelah diambil data di pool kosong
			// kalo goroutine gadapet data nanti akan direturn nilai nil, karena datanya belum balik
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
