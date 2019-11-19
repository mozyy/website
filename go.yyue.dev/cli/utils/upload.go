package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
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
func UploadDeploy(filePath, dest string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", filepath.Base(filePath))
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, dest, body)
	if err != nil {
		return err
	}
	// 添加浏览器 User Agent 模拟电脑chrome
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	userCookie := &http.Cookie{Name: "fb_user", Value: url.QueryEscape("周志强"), Expires: time.Now().Add(time.Hour)}
	request.AddCookie(userCookie)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	alert := doc.Find("script").Eq(0).Text() //alert("上传成功")
	reg := regexp.MustCompile(`alert\(\"(\S+)\"\)`)

	result := reg.FindAllStringSubmatch(alert, -1)
	if len(result) > 0 && result[0][1] == "上传成功" {
		fmt.Println("上传成功")
	} else {
		return fmt.Errorf(doc.Text())
	}
	return nil
}
