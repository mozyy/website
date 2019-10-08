package fetcher

import (
	"strings"
	"sync"
	"testing"
)

func TestFetch(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			b := Fetch("https://www.baidu.com/")
			if !strings.HasPrefix(string(b), "HTTP/1.1 200 OK") {
				t.Errorf("first 15 strings want:\n%s, got:\n%s", string(b[:15]), "HTTP/1.1 200 OK")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
