package goroutinepzngo

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// cara pakenya tinggal tambah `go` sebelum fungsi yang akan dipanggil
	go RunHelloWorld()
	// *goroutine kurang cocok kalo dipake di fungsi yang mengembalikan nilai, karena nilainya gaakan bisa ditangkap
	// goroutine = async
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
