package goroutinepzngo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait() // kalo ada cond.Wait() goroutine harus nunggu sampe ada cond.Signal()/cond.Broadcast() biar bisa lanjut, kalo gaada akan deadlock karena goroutine lain harus nunggu
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal() // cond.Signal() akan dijalanin 1-1 tiap 1 detik
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast() // cuman nunggu 1 detik, setelah diunlock semua bakal jalan, mirip locking biasa
	}()

	group.Wait()
}
