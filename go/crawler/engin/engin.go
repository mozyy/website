package engin

import "crawler/fetcher"

// Run is the entry point for engin
func (request *Request) Run() {
	fetcher.Fetch(request.URL)
}
