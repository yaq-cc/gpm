package cmd

import (
	"encoding/json"
	"fmt"
	"io"
)

type Workspace struct {
	Folders  []*Folder `json:"folders"`
	Settings *Settings `json:"settings"`
}

type Folder struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type Settings struct {
	Gopls *Gopls `json:"gopls"`
}

type Gopls struct {
	BuildDirectoryFilters []string `json:"build.directoryFilters"`
}

func (ws *Workspace) Create(w io.Writer, projectName string) error {
	ws.Folders = make([]*Folder, 0)
	folder := &Folder{
		Path: ".",
		Name: fmt.Sprintf("%s [ROOT]", projectName),
	}
	buildDirFilters := make([]string, 0)
	buildDirFilters = append(buildDirFilters, "-../../")

	ws.Folders = append(ws.Folders, folder)
	ws.Settings = &Settings{
		Gopls: &Gopls{
			BuildDirectoryFilters: append(buildDirFilters, "+"+projectName),
		},
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	err := enc.Encode(ws)
	if err != nil {
		return err
	}
	return nil
}
