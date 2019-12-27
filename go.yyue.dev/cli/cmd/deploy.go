package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"go.yyue.dev/cli/util"
)

var (
	build       string
	packageName string
	uploadURL   string

	deployCmd = &cobra.Command{
		Use:   "dep",
		Short: "deploy project to server",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires base path")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			// basePath := "d://work/video-session/"
			// build := "../build"
			// packageName := "sss"
			// uploadURL := "http://192.168.88.122:9990/upfile.php?pk_id=202"

			basePath := args[0]

			temp, err := ioutil.TempDir(basePath, "build")
			if err != nil {
				fmt.Println("make temp dir error: ", err)
				return
			}
			defer os.RemoveAll(temp)
			dist, err := util.Compress(temp, filepath.Join("../", build), packageName)
			if err != nil {
				fmt.Println("compress error: ", err)
				return
			}
			err = util.UploadDeploy(dist, uploadURL)
			if err != nil {
				fmt.Println("upload error: ", err)
				return
			}
			fmt.Println("部署成功")
		},
	}
)

func init() {
	deployCmd.Flags().StringVarP(&build, "build", "b", "", "build dire")
	deployCmd.MarkFlagRequired("build")
	deployCmd.Flags().StringVarP(&packageName, "package", "p", "", "package name")
	deployCmd.MarkFlagRequired("package")
	deployCmd.Flags().StringVarP(&uploadURL, "upload", "u", "", "upload URL")
	deployCmd.MarkFlagRequired("upload")
}
