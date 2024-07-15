package goroutinepzngo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // fungsi Done() sama kayak Add(-1)
	// setelah jalan, jangan lupa dikasih Done() biar ga deadlock

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait() // Wait() akan selesai ketika Add(0), cara biar Add(0) adalah pake Done() setiap selesai 1 proses yang dijalanin
	fmt.Println("Complete")
}
