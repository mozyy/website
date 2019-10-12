package fetcher

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"yyue.dev/common/utils"
)

var ticker = time.NewTicker(10 * time.Millisecond)

// Fetch url
func Fetch(url string) *goquery.Document {
	// 每次请求间隔一段时间
	<-ticker.C
	log.Printf("Fetcher: fetching url: %s", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	utils.PanicErr(err)
	// 添加浏览器 User Agent 模拟电脑chrome
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	client := http.Client{}
	resp, err := client.Do(request)
	utils.PanicErr(err)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		utils.PanicErr(fmt.Errorf("Resp status fail, code is : %d", resp.StatusCode))
	}
	// response, err := httputil.DumpResponse(resp, true)
	// fileName := "index.html"
	// _, err = os.Stat(fileName)
	// var f *os.File
	// if os.IsNotExist(err) {
	// 	f, err = os.Create(fileName)
	// } else {
	// 	f, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	// }
	// utils.PanicErr(err)
	// defer f.Close()
	// _, err = f.Write(response)
	// utils.PanicErr(err)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	utils.PanicErr(err)
	return doc
}
