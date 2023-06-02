# GPM: Go Project Manager

The Go Project Manager (gpm) is an opinionated tool for scaffolding and quickstarting Go based projects in GCP.  

# Features

- Projects created with `gpm` include:
- cmd and pkg folders,
- An activation script (activate) that maintains key environment variables used by scripts and other deployment tools.
- A `cloudbuild.yaml` file,  dynamically generated based on your project's data for deployment to Cloud Run.
- A code-workspace file for VS Code and Code* IDE derivatives like Code-Server, Theia Eclipse, and others.  
- A basic Dockerfile for building and running a Go binary.
- A `scripts.yaml` file that maintains commonly used scripts for docker and git.  The `scripts.yaml` file is wholly extensible and you can add / remove scripts as you'd like. (more on this later!)
- `gpm` also has some basic code generation capabilities:
- Cloud provides the infrastructure to run a reliable HTTP(S) endpoint; as such gpm provides some basic boilerplate code generation.  A reverse proxy template is also provided.
- More common or tricky patterns will be provided as time permits.  We're looking at making this configurable and user extensible.  

# Usage

## Creating a project
gpm projects create `project`

## Running a script
gpm scripts run `script` OR
gpx `script` 

See the scripts file.

## Generating boilerplate code
gpm code generate `code` /path/`code.go` 
two 'helpers' available: 
- server (creates HTTP server), 
- reverse-proxy (tempalte for reverse proxies)

