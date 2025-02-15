package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B)  {
	b.Run("Fahril", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fahril")
		}
	})
	b.Run("Hadi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hadi")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fahril")
	}
}

func BenchmarkHelloWorldHadi(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hadi")
	}
}

func TestTableHelloWorld(t *testing.T)  {
	tests := []struct {
		name	 string
		request  string
		expected string
	}{
		{
			name:	"Fahril",
			request: "Fahril",
			expected: "Hello Fahril",
		},
		{
			name:	"Hadi",
			request: "Hadi",
			expected: "Hello Hadi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T)  {
	t.Run("Fahril", func(t *testing.T)  {
		result := HelloWorld("Fahril")
		require.Equal(t, "Hello Fahril", result, "Result must be 'Hello Fahril'")
	})
	t.Run("Hadi", func(t *testing.T) {
		result := HelloWorld("Hadi")
		require.Equal(t, "Hello Hadi", result, "Result must be 'Hello Hadi'")
	})
}

func TestMain(m *testing.M)  {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Fahril")
	require.Equal(t, "Hello Fahril", result, "Result must be 'Hello Fahril'")
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Fahril")
	require.Equal(t, "Hello Fahril", result, "Result must be 'Hello Fahril'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("Fahril")
	assert.Equal(t, "Hello Fahril", result, "Result must be 'Hello Fahril'")
	fmt.Println("TestHelloWorld with Assert Done")
}

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