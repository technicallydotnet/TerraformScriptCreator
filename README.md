## Overview
A simple application for creating Terraform configs for Azure via a CLI application.

## How to use
1. Compile the application (if the _executable_ in the repository is not being used) with `go build TerraformScriptCreator.go`
2. Run the _executable_ from a console, specificing the following parameters:
 - `vmNumber` (required) - The number of virtual machines to be created. Default value is _0_.
 - `webAppNumber` (required) - The number of web apps to be created. Each web app will run on it's own app service plan. Default value is _0_.
 - `acrNumber` (required) - The number of Azure Container Registries to be created. Default value is _0_.
 - `aksNumber` (required) - The number of Azure Kubernetes clusters to be created. Default value is _0_.
 - `os` (required when creating virtual machines) - operating system of virtual machines. Default value is _windows_.
 - `location` (optional) - location of resources and resource group. Default value is _West Europe_.
 - `spec` (optional) - specification of resources to be created. By default, there are 3 categories: _low_, _medium_ and _high_. Default value is _low_.

 **NOTE** - If no parameters are entered, no resources will be created. One of the following must have a value set: _vmNumber_, _webAppNumber_, _acrNumber_ or _aksNumber_

## Future features
- Being able to set spec for each resource category with an external text file
- Creating/deleting different specification levels with an external file
- Adding an optional parameter, which will add optional Terraform config to resources
