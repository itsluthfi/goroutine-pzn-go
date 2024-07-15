package goroutinepzngo

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCrateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // setelah digunakan wajib dihapus

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Luthfi Izzuddin Hanif" // ngirim data ke channel, by default channel cuman bisa nampung 1 data
		// go routine bakal nunggu ada go routine lain yg ambil data, kalo gaada bakal ditunggu terus sampe ada, proses stuck/keblock/deadlock
		// jadi pastiin kalo ngirim data ke channel ada yang nerima
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel // ambil data dari channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	// channel kalo ngirim value as params dia pass by reference jadi ga perlu pake pointer
	time.Sleep(2 * time.Second)
	channel <- "Luthfi Izzuddin Hanif"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // setelah digunakan wajib dihapus

	go GiveMeResponse(channel)

	data := <-channel // ambil data dari channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	// chan<- berarti mengirim data ke channel
	time.Sleep(2 * time.Second)
	channel <- "Luthfi Izzuddin Hanif"
}

func OnlyOut(channel <-chan string) {
	// <-chan berarti mengambil data dari channel/data keluar dari channel
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // setelah digunakan wajib dihapus

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // buffer didefiniskan di integer terakhir
	// kalo 3 berarti bisa 3 data yang dimasukin ke buffer
	defer close(channel)

	go func() {
		channel <- "Luthfi"
		channel <- "Izzuddin"
		channel <- "Hanif"
	}()

	// data di buffer ga diambil tetep kesimpen, tapi kalo udah penuh dan mau ngirim, data yang baru nunggu ada data di buffer ada yang diambil dulu

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	// dipake ketika gatau berapa banyak data yang bakal dikirim/dimasukin ke channel
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel) // pastikan diclose setelah dipakai, kalo ga nanti infine loop dan test/proses ga selesai2/deadlock
	}()

	// for range akan ngeblock di satu channel aja
	for data := range channel { // perulangan akan berhenti kalo channelnya udah diclose
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
