package main

import (
	"github.com/spf13/cobra"

	"github.com/stackup-wallet/stackup-bundler/internal/start"
)

var rootCmd = &cobra.Command{
	Use:   "stackup-bundler",
	Short: "ERC-4337 Bundler",
	Long:  "A modular Go implementation of an ERC-4337 Bundler.",
}

func main() {
	//err := rootCmd.Execute()
	//if err != nil {
	//	os.Exit(1)
	//}

	start.PrivateMode()

	//start.SearcherMode()
}

func init() {}
