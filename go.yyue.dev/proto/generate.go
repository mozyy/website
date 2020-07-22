package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"go.yyue.dev/common/utils"
)

//go:generate go run $GOFILE

const (
	goOutPath  = "../../"
	webOutPath = "../../www/client/proto"
)

func main() {
	dir, err := os.Getwd()
	utils.PanicErr(err)
	err = filepath.Walk("./", func (p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
		}
		fmt.Println(p, info.IsDir(), filepath.Ext(p))
		if (!info.IsDir() && filepath.Ext(p) == ".proto") {

			args := []string{
				"-I=.",
				// 分开写会生成两个文件, 下面的写一起只生成一个文件 ,
				// 下面的import的context是golang.org/x/net/context, 以被内部context取代,
				// 所以用上面的
				fmt.Sprintf("--micro_out=%s", goOutPath),
				fmt.Sprintf("--go_out=%s", goOutPath),
				// fmt.Sprintf("--go_out=plugins=micro:%s", goOutPath),
				fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", webOutPath),
				fmt.Sprintf("--grpc-web_out=import_style=typescript,mode=grpcwebtext:%s", webOutPath),
				p,
			}
			cmd := exec.Command("protoc", args...)
			cmd.Dir = dir
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("protoc [%s] error: %s, out: \n%s\n", p, err, out)
			} else {
				fmt.Printf("protoc success: %s\n", p)
			}
		}
    return nil
	})
	utils.PanicErr(err)
}