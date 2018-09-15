# Papercall Formatter

Parse Papercall Downloaded submissions into Speaker / Program markdown files for [DevOpsDays-Web](https://github.com/devopsdays/devopsdays-web)

Works for following json format:

```json
$ cat testdata/download.json | jq '.[0] | keys '
[
  "abstract",
  "additional_info",
  "audience_level",
  "avatar",
  "bio",
  "confirmed",
  "created_at",
  "description",
  "email",
  "location",
  "name",
  "notes",
  "organization",
  "rating",
  "shirt_size",
  "state",
  "tags",
  "talk_format",
  "title",
  "twitter",
  "url"
]
```

## Make Targets

```bash
help                           List make targets & descriptions

build                          Build and print help output
clean                          Clean generated files
program                        Clear and re-generate program files based on accepted submissions
speakers                       Clear and re-generate speakers profiles based on accepted submissions
```

## CLI Usage

```bash
NAME:
   Papercall Format - Parse Papercall submissions from json

USAGE:
   papercallfmt [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   so0k

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level value                      Log level (panic, fatal, error, warn, info, or debug) (default: "error") [$LOG_LEVEL]
   --source json, -s json                 Source json (default: "download.json")
   --destination directory, -d directory  Destination directory to render in - must exist (default: "output/speakers/")
   --template value, -t value             Desired template used to render output (default: "templates/speaker.md.tpl")
   --state-filter value, -f value         Filter to only submissions of this state (default: "accepted")
   --help, -h                             show help
   --version, -v                          print the version
```
