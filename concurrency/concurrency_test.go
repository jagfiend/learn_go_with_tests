package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// test
func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestWebsiteChecker(t *testing.T) {
	t.Run("returns checks as expected", func(t *testing.T) {
		websites := []string{
			"https://google.com",
			"https://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		want := map[string]bool{
			"https://google.com":          true,
			"https://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":     false,
		}

		got := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

// benchmark
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "earl?"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
