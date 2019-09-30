package fetcher

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"yyue.dev/common/utils"
)

// Fetch url
func Fetch(url string) []byte {
	log.Printf("Fetcher: fetching url: %s", url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	utils.PanicErr(err)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	client := http.Client{}
	resp, err := client.Do(request)
	utils.PanicErr(err)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		utils.PanicErr(fmt.Errorf("Resp status fail, code is : %d", resp.StatusCode))
	}
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		utils.PanicErr(err)
	}
	return response
}
