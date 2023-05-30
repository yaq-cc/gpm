/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

		var dirs string
		var fname string
		var reader *strings.Reader

		switch template := args[0]; template {
		case "server":
			fname = "server.go"
			reader = SourceCodeHTTPServer
		case "reverse-proxy":
			fname = "proxy.go"
			reader = SourceCodeReverseProxy
		default:
			log.Fatal("unrecognized template")
		}

		if len(args) == 2 {
			dirs, _ = filepath.Split(args[1])
			fname = args[1]
		}

		if err := os.MkdirAll(dirs, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		tFile, err := os.Create(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer tFile.Close()

		_, err = io.Copy(tFile, reader)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	codeCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
