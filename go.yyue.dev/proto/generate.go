package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"go.yyue.dev/common/utils"
)

//go:generate go run $GOFILE

const (
	goOutPath  = "../../"
	webOutPath = "../web/"
)

func main() {
	files, err := readDir("./")
	utils.PanicErr(err)
	for _, file := range files {
		if strings.HasSuffix(file, ".proto") {
			dir, err := os.Getwd()
			utils.PanicErr(err)
			args := []string{
				"-I=.",
				fmt.Sprintf("--go_out=plugins=micro:%s", goOutPath),
				fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", webOutPath),
				fmt.Sprintf("--grpc-web_out=import_style=typescript,mode=grpcwebtext:%s", webOutPath),
				file,
			}
			cmd := exec.Command("protoc", args...)
			cmd.Dir = dir
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("protoc error: %s, out: %s\n", err, out)
			} else {
				fmt.Printf("protoc success: %s\n", file)
			}
		}
	}
}

func readDir(dirname string) ([]string, error) {
	readFiles, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, readFile := range readFiles {
		// fmt.Println(readFile.Name())
		if readFile.IsDir() {
			subFiles, err := readDir(dirname + readFile.Name() + "/")
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			files = append(files, dirname+readFile.Name())
		}
	}
	return files, nil
}
