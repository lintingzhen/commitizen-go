# Commitizen-go
> Command line utility to standardize git commit messages, golang version.

![demo](docs/images/demo.gif)

Contents
- Getting Started
- Usage
- Configure

## Getting Started
### installation with Homebrew:
```
$ brew tap lintingzhen/tap
$ brew install commitizen-go
$ sudo commitizen --install
```
### installation with AUR package (Arch Linux):
```
$ yay commitizen-go
```
### installation with source code:
```
$ make && make install
```
or 
```
$ make && ./commitizen-go --install
```
### commit with commitizen
```
$ git cz
```

## Usage
```
Usage:
  commitizen-go [flags]
  commitizen-go [command]

Available Commands:
  help        Help about any command
  install     Install this tool to git-core as git-cz
  version     Print version information and quit

Flags:
  -a, --all     tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected
  -d, --debug   debug mode, output debug info to debug.log
  -h, --help    help for commitizen-go

Use "commitizen-go [command] --help" for more information about a command.
```

## Configure
You can customize your commit type, items, and how to assemble these items through the configuration file.
Configuration file path is `~/.git-czrc`, and the format is the same as the defaultConfig string in the file commit/defaultConfig.go.
