package helper

import "testing"

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