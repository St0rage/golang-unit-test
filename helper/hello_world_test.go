package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Dani Yudistira Maulana",
			request: "Dani Yudistira Maulana",
		},
		{
			name:    "Veronica Christine Ningtyas",
			request: "Veronica Christine Ningtyas",
		},
	}

	for _, bebenchmark := range benchmarks {
		b.Run(bebenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bebenchmark.request)
			}
		})
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Dani", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dani")
		}
	})
	b.Run("Veronica", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Veronica")
		}
	})
}

// Benchmark
func BenchmarkHelloWorldDani(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dani")
	}
}

func BenchmarkHelloWorldVeronica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Veronica")
	}
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Dani",
			request:  "Dani",
			expected: "Hello Dani",
		},
		{
			name:     "Dian",
			request:  "Dian",
			expected: "Hello Dian",
		},
		{
			name:     "Veronica",
			request:  "Veronica",
			expected: "Hello Veronica",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Dani", func(t *testing.T) {
		result := HelloWorld("Dani")
		require.Equal(t, "Hello Dani", result, "Result must be 'Hello Dani'")
	})
	t.Run("Veronica", func(t *testing.T) {
		result := HelloWorld("Veronica")
		require.Equal(t, "Hello Veronica", result, "Result must be 'Hello Veronica'")
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Dani")
	require.Equal(t, "Hello Dani", result, "Result must be 'Hello Dani'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dani")
	require.Equal(t, "Hello Dani", result, "Result must be 'Hello Dani'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Dani")
	assert.Equal(t, "Hello Dani", result, "Result must be 'Hello Dani'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldDani(t *testing.T) {
	result := HelloWorld("Dani")

	if result != "Hello Dani" {
		// error
		t.Error("Result must be 'Hello Dani'")
	}

	fmt.Println("TestHelloWorldDani Done")
}

func TestHelloWorldVeronica(t *testing.T) {
	result := HelloWorld("Veronica")

	if result != "Hello Veronica" {
		// error
		t.Fatal("Result must be 'Hello Veronica'")
	}

	fmt.Println("TestHelloWorldVeronica Done")
}
