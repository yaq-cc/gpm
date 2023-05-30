package cmd

import (
	_ "embed"
	"strings"
)

//go:embed templates/Dockerfile
var dockerfile string

var Dockerfile = strings.NewReader(dockerfile)

//go:embed templates/script-git-repo-quickstart.sh
var gitQuickstartScript string

//go:embed templates/script-git-quick-push.sh
var gitQuickPushScript string

//go:embed templates/server.go
var sourceCodeHTTPServer string

var SourceCodeHTTPServer = strings.NewReader(sourceCodeHTTPServer)

//go:embed templates/reverseProxy.go
var sourceCodeReverseProxy string

var SourceCodeReverseProxy = strings.NewReader(sourceCodeReverseProxy)
