package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"time"
)

var (
	devSrc   = "192.168.2.252"
	user     = "root"
	password = "shulaibao"
)

// UploadDev upload local directory to linux directory
func UploadDev(client, dest string) error {
	cmd := exec.Command("bash", "")
	cmdWriter, _ := cmd.StdinPipe()
	cmd.Start()

	cmdString := fmt.Sprintf("sshpass -p %s scp -r %s %s@%s:%s", password, devSrc, user, client, dest)

	cmdWriter.Write([]byte(cmdString + "\n"))
	cmdWriter.Write([]byte("exit" + "\n"))

	cmd.Wait()
	return nil
}

// UploadDeploy file to deploy server
func UploadDeploy(dir, dest string) error {

	request, err := http.NewRequest(http.MethodPost, dest, nil)
	if err != nil {
		return err
	}
	// 添加浏览器 User Agent 模拟电脑chrome
	// request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	userCookie := &http.Cookie{Name: "fb_user", Value: url.QueryEscape("周志强"), Expires: time.Now().Add(time.Hour)}
	request.AddCookie(userCookie)
	request.Header.Set("Content-Type", "multipart/form-data")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	str := string(response)
	fmt.Println(str)
	return nil
}
