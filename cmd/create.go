/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		dirName := args[0]
		fmt.Println("create called")
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				log.Fatalf("gpm projects create failed: could not create \"%s\" as it already exists", dirName)
			}
			log.Fatal(err)
		}
		cmdDir := dirName + "/cmd"
		pkgDir := dirName + "/pkg"
		err = os.MkdirAll(cmdDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err = os.MkdirAll(pkgDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		wsName := dirName + "/" + dirName + ".code-workspace"
		file, err := os.Create(wsName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		err = (&Workspace{}).Create(file, dirName)
		if err != nil {
			log.Fatal(err)
		}

		dfName := dirName + "/Dockerfile"
		dfFile, err := os.Create(dfName)
		if err != nil {
			log.Fatal(err)
		}
		defer dfFile.Close()
		_, err = io.Copy(dfFile, Dockerfile)
		if err != nil {
			log.Fatal(err)
		}

		cbName := dirName + "/cloudbuild.yaml"
		cbFile, err := os.Create(cbName)
		if err != nil {
			log.Fatal(err)
		}
		defer cbFile.Close()

		err = (&CloudBuild{}).Create(cbFile, dirName)
		if err != nil {
			log.Fatal(err)
		}

		scriptsName := dirName + "/scripts.yaml"
		sFile, err := os.Create(scriptsName)
		if err != nil {
			log.Fatal(err)
		}
		err = Scripts.Create(sFile, dirName)
		if err != nil {
			log.Fatal(err)
		}

		// Preconfigured environment variables.
		// Decouples the data from the scripts - change configuration data,
		// the scripts will always point to these environment variables
		cfgName := dirName + "/activate"
		cfgFile, err := os.Create(cfgName)
		if err != nil {
			log.Fatal(err)
		}
		defer cfgFile.Close()
		cfgFile.WriteString("echo 'Thank you for using gpm, you're environment is almost ready..'")
		cfgFile.WriteString("export GOPATH=/develop/go\n")
		cfgFile.WriteString("export PATH=$PATH:$GOPATH/bin\n")
		cfgFile.WriteString(fmt.Sprintf("export PROJECT_NAME=%s\n", dirName))
		cfgFile.WriteString("export INIT_PS1=$PS1\n")
		cfgFile.WriteString("export PS1=\"($PROJECT_NAME) $INIT_PS1\"\n")
		cfgFile.WriteString("export GH_USER_NAME=YvanJAquino\n")
		cfgFile.WriteString("export GH_AUTHOR_NAME=\"Yvan J. Aquino\"\n")
		cfgFile.WriteString("export GH_USER_EMAIL=yaquino@google.com\n")
		cfgFile.WriteString("export GH_PA_TOKEN=MY_PAT\n")
		cfgFile.WriteString("export GH_REPO_URL=https://$GH_PA_TOKEN@github.com/$GH_USER_NAME/$PROJECT_NAME.git\n")
		cfgFile.WriteString("export GCP_PROJECT_ID=holy-diver-297719\n")
		cfgFile.WriteString("export GCP_AR_FOLDER=containers\n")
		cfgFile.WriteString("export GCP_REGION=us-central1\n")
		cfgFile.WriteString("echo 'Replace `gpm scripts run` with `gpx` (an alias)'\n")
		cfgFile.WriteString("alias gpx=\"gpm scripts run\"\n")
		cfgFile.WriteString("deactivate () {\n\tPS1=$INIT_PS1\n\texport PS1\n\tunset -f deactivate\n}\n")
		cfgFile.WriteString("echo 'Return to your regular environment with 'deactivate'")

	},
}

func init() {
	projectsCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
