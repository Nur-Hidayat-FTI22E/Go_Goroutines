package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Nur Hidayat"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	// channel <- "Dayat"

	// fmt.Println(<-channel)

	// close(channel)
}
