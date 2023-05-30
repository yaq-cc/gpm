package cmd

import (
	"io"

	"gopkg.in/yaml.v3"
)

var Scripts = scripts{}

type scripts map[string]any

func (s scripts) Create(w io.Writer, projectName string) error {
	s = scripts{
		"go": scripts{
			"mod": scripts{
				"init": "go mod init github.com/$GH_USER_NAME/$PROJECT_NAME",
			},
		},
		"docker": scripts{
			"build": "docker build -t local.$PROJECT_NAME .",
			"run": scripts{
				"default": "docker run -it --rm --network=host local.$PROJECT_NAME",
				"shell":   "docker run -it --rm --network=host local.$PROJECT_NAME /bin/bash",
			},
		},
		"git": scripts{
			"repo-quickstart": gitQuickstartScript,
			"quick-push":      gitQuickPushScript,
		},
	}
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	err := enc.Encode(s)
	if err != nil {
		return err
	}
	return nil
}
