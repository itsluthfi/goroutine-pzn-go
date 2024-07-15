package goroutinepzngo

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)      // kalo ada warning diignore aja karena setelahnya ada once.Do() jadi gaakan ada race condition
			once.Do(OnlyOnce) // once.Do() cuman bisa dipake buat fungsi yg gaada parameternya
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter ", counter)
}
