package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "ttd",
		Short: "ttd is a totodi tools",
	}
)

func Execute() {
	// err := doc.GenMarkdownTree(rootCmd, "./")
	// if err != nil {
	// 	fmt.Println("gen doc error: ", err)
	// }
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
