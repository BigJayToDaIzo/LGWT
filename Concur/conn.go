package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	// send out all the jobs into the results channel
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// patiently await the completion of the resultChannel returning data
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel       // grab FIFO result from the queue
		results[r.string] = r.bool // unpack result into results map
	}
	return results
}
