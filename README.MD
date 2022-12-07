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
$ sudo commitizen-go install
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
$ make && ./commitizen-go install
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
You can set configuration file that `.git-czrc` at repository root or home directory. (You can also add the extension to file, like `.git-czrc.json`)
The configure file that located in repository root have a priority over the one in home directory.
The format is the same as the defaultConfig string in the file [commit/defaultConfig.go](https://github.com/lintingzhen/commitizen-go/blob/master/commit/defaultConfig.go).

Type item like that:
```
...
{
	"name": "type",
	"desc": "Select the type of change that you're committing:",
	"form": "select",
	"options": [
		{ "name": "feat", "desc": "feat: A new feature" },
		{ "name": "fix", "desc": "fix: A bug fix" },
		{ "name": "docs", "desc": "docs: Documentation only changes" },
		{
		  "name": "style",
		  "desc":
			"style: Changes that do not affect the meaning of the code\n            (white-space, formatting, missing semi-colons, etc)"
		},
		{
		  "name": "refactor",
		  "desc": "refactor: A code change that neither fixes a bug nor adds a feature"
		},
		{
		  "name": "perf",
		  "desc": "perf: A code change that improves performance"
		},
		{ "name": "test", "desc": "test: Adding missing tests" },
		{
		  "name": "chore",
		  "desc":
			"chore: Changes to the build process or auxiliary tools\n            and libraries such as documentation generation"
		},
		{ "name": "revert", "desc": "revert: Revert to a commit" },
		{ "name": "WIP", "desc": "WIP: Work in progress" }
	],
	"required": true
},
...
```
Template like that:
```
...
"template": "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}"
...
```
