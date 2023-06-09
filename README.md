# Support for Conjur Vault integration added.
This version provides a new backend to the argocd vault plugin. We have tested it for argocd 2.6.4
# How to use it:
An image is already available on: https://hub.docker.com/r/itdistrict/argocd-vault-plugin

You can also build your own image:
1) clone/download this repo.
2) build the software by using: GOOS=linux GOARCH=amd64 go build -o argocd-vault-plugin 
3) build the docker image with the provided Dockerfile
4) tag the image and upload it to your repository.
5) use it :-)

# What's new?
We have added conjur vault support to argocd-vault-plugin. This plugin provides two ways of authentication. 
1) authentication by using the sidecar authentication (preferred)
2) authentication by using a static api-key

# What's next?
We will add another project to show the usage of this integration and will work on a pull request into the original project.

# argocd-vault-plugin
![Pipeline](https://github.com/argoproj-labs/argocd-vault-plugin/workflows/Pipeline/badge.svg)
![Code Scanning](https://github.com/argoproj-labs/argocd-vault-plugin/workflows/Code%20Scanning/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/argoproj-labs/argocd-vault-plugin)](https://goreportcard.com/report/github.com/argoproj-labs/argocd-vault-plugin)
![Downloads](https://img.shields.io/github/downloads/IBM/argocd-vault-plugin/total?logo=github)
[![codecov](https://codecov.io/gh/argoproj-labs/argocd-vault-plugin/branch/main/graph/badge.svg?token=6Xr7V8AMTE)](https://codecov.io/gh/argoproj-labs/argocd-vault-plugin)

<img src="https://github.com/argoproj-labs/argocd-vault-plugin/raw/main/assets/argo_vault_logo.png" width="300">

An Argo CD plugin to retrieve secrets from various Secret Management tools (HashiCorp Vault, IBM Cloud Secrets Manager, AWS Secrets Manager, etc.) and inject them into Kubernetes resources

### Why use this plugin?
This plugin is aimed at helping to solve the issue of secret management with GitOps and Argo CD. We wanted to find a simple way to utilize Secret Management tools without having to rely on an operator or custom resource definition. This plugin can be used not just for secrets but also for deployments, configMaps or any other Kubernetes resource.

## Documentation
You can our full set of documentation at https://argocd-vault-plugin.readthedocs.io/

## Contributing
Interested in contributing? Please read our contributing documentation [here](./CONTRIBUTING.md) to get started!

## Blogs
- [Solving ArgoCD Secret Management with the argocd-vault-plugin](https://itnext.io/argocd-secret-management-with-argocd-vault-plugin-539f104aff05)
- [Introducing argocd-vault-plugin v1.0!](https://itnext.io/introducing-argocd-vault-plugin-v1-0-708433294b2d)
- [How to Use HashiCorp Vault and Argo CD for GitOps on OpenShift](https://cloud.redhat.com/blog/how-to-use-hashicorp-vault-and-argo-cd-for-gitops-on-openshift)

## Presentations
- [Shh, It’s a Secret: Managing Your Secrets in a GitOps Way - Jake Wernette & Josh Kayani, IBM](https://youtu.be/7L6nSuKbC2c)
