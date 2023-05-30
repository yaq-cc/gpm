/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Open and decode the scripts.yaml file.
		userScripts := scripts{}
		file, err := os.Open("scripts.yaml")
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.NewDecoder(file).Decode(userScripts)
		if err != nil {
			log.Fatal(err)
		}

		// Walk the scripts file, looking for the script itself.
		var script string
		scriptMap := map[string]any(userScripts)
		for _, arg := range args {
			scriptValue, ok := scriptMap[arg]
			if !ok {
				log.Fatal("unrecognized script")
			}
			switch svt := scriptValue.(type) {
			case scripts:
				scriptMap = map[string]any(svt)
				continue
			case string:
				script = svt
			default:
				log.Fatal("scripts and script groups must be in dictionary format, do not use arrays and lists!", svt)
			}
		}

		// Execute the script.
		err = executeScript(script)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	scriptsCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func executeScript(script string) error {
	bin := exec.Command("/bin/bash", "-c", script)
	bin.Stdout = os.Stdout
	bin.Stdin = os.Stdin
	bin.Stderr = os.Stderr
	if err := bin.Run(); err != nil {
		return err
	}
	return nil
}
