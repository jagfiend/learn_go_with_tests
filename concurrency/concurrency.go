package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// make a channel to manage results of gorountines below and prevent race conditions
	resultChannel := make(chan result)

	// loop over urls and process them concurrenctly with goroutines
	for _, url := range urls {
		go func(u string) {
			// <- sends the result struct into the channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	results := make(map[string]bool)

	// assign the result struct out of the channel in loop order
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
