package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Fahril")

	if result != "Hello Fahril" {
		// error
		panic("Result is not 'Hello Fahril'")
	}

	// gunakan perintah go test (menjalankan unit test)
	// gunakan perintah go test -v (melihat detail function test apa saja yang sudah di running)
	// gunakan perintah go test -v -run TestNamaFunction (memilih function unit test mana yang ingin di running)
	// gunakan perintah go test -v ./... (menjalankan semua unit test dari top folder module nya)
}

func TestHelloWorldHadi(t *testing.T)  {
	result := HelloWorld("Hadi")

	if result != "Hello Hadi" {
		// error
		t.Error("Result must be 'Hello Hadi'")
	}

	fmt.Println("TestHelloWorldHadi Done")
}

func TestHelloWorldAbu(t *testing.T)  {
	result := HelloWorld("Abu")

	if result != "Hello Abu" {
		// error
		t.Fatal("Result must be 'Hello Abu'")
	}

	fmt.Println("TestHelloWorldAbu Done")
}