package cmd

import (
	"io"

	"gopkg.in/yaml.v3"
)

type CloudBuild struct {
	Steps         []*Step        `yaml:"steps"`
	Substitutions *Substitutions `yaml:"substitutions"`
}
type Step struct {
	ID         string   `yaml:"id"`
	WaitFor    []string `yaml:"waitFor,omitempty"`
	Name       string   `yaml:"name"`
	Dir        *string  `yaml:"dir,omitempty"`
	Entrypoint string   `yaml:"entrypoint"`
	Args       []string `yaml:"args"`
}
type Substitutions struct {
	Service string `yaml:"_SERVICE"`
	Region  string `yaml:"_REGION"`
}

func (cb *CloudBuild) Create(w io.Writer, projectName string) error {
	cb.Steps = []*Step{
		{
			ID:         "docker-build-push",
			WaitFor:    []string{"-"},
			Name:       "gcr.io/cloud-builders/docker",
			Entrypoint: "bash",
			Args: []string{
				"-c",
				"docker build -t gcr.io/$PROJECT_ID/${_SERVICE} . &&\ndocker push gcr.io/$PROJECT_ID/${_SERVICE}",
			},
		}, {
			ID:         "gcloud-run-deploy",
			WaitFor:    []string{"docker-build-push"},
			Name:       "gcr.io/google.com/cloudsdktool/cloud-sdk",
			Entrypoint: "bash",
			Args: []string{
				"-c",
				"gcloud run deploy ${_SERVICE} \\\n--project $PROJECT_ID \\\n--image gcr.io/$PROJECT_ID/${_SERVICE} \\\n--timeout 5m \\\n--region ${_REGION} \\\n--no-cpu-throttling \\\n--min-instances 0 \\\n--max-instances 5 \\\n--allow-unauthenticated",
			},
		},
	}
	cb.Substitutions = &Substitutions{
		Service: projectName,
		Region:  "us-central1",
	}

	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	err := enc.Encode(cb)
	if err != nil {
		return err
	}
	return nil
}
