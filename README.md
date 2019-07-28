# mageXT [![Build Status](https://travis-ci.org/pvormste/magext.svg?branch=master)](https://travis-ci.org/pvormste/magext)

mageXT contains some extensions for usage in go [magefiles](https://github.com/magefile/mage).

You can see the extensions in this repository in action:
 - `mage.go`: is the mage 'executable'
 - `mage_targets.go`: contains the actual targets for mage an uses packages like `gocmd`, `cilintcmd`, etc.
 - `mage_env.go`: defines the mage environment and uses the env names from the packages

## Extensions

| package | description | godoc |
| ------- | ----------- | ----- |
| gocmd | go commands like `get`, `install` | [godoc](https://godoc.org/github.com/pvormste/magext/gocmd) |
| cilintcmd | commands for golangci-lint | [godoc](https://godoc.org/github.com/pvormste/magext/cilintcmd) |