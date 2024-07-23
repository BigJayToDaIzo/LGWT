package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc func(url string) bool, sites []string) map[string]bool {
	siteMap := make(map[string]bool)
	siteChannel := make(chan result)
	defer close(siteChannel)
	for _, site := range sites {
		go func(s string) {
			siteChannel <- result{s, wc(s)}
		}(site)
	}
	for i := 0; i < len(sites); i++ {
		site := <-siteChannel
		siteMap[site.string] = site.bool
	}
	return siteMap
}
