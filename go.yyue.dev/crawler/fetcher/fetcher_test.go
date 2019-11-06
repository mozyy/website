package fetcher

import (
	"sync"
	"testing"
)

func TestFetch(t *testing.T) {
	wg := sync.WaitGroup{}
	title := "百度一下，你就知道"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			b := Fetch("https://www.baidu.com/")
			if b.Find("title").Text() != title {
				t.Errorf("first 15 strings want:\n%s, got:\n%s", title, b.Find("title").Text())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
