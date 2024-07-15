package goroutinepzngo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1) // ditambahin di sebelum goroutine
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1) // atomic dipake buat manipulasi data primitive di goroutine/proses concurrent biar gausah pake mutex
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
