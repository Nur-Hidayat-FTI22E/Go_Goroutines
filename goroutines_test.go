package gogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWord() {
	fmt.Println("Hello Word")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWord()

	fmt.Println("Ups")

	time.Sleep(10 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
