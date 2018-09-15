JSON_FILE ?= "testdata/test.json"
LOG_LEVEL ?= "error"
BUILD_NUMBER ?= 0

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')


help: ## List make targets & descriptions
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

bin/papercallfmt: $(SOURCES)
	go build -ldflags "-X main.build=${BUILD_NUMBER}" -o bin/papercallfmt

speakers: bin/papercallfmt  ## Clear and re-generate speakers profiles based on accepted submissions (set JSON_FILE / LOG_LEVEL)
	-rm output/speakers/*.md
	bin/papercallfmt -s $(JSON_FILE) -t templates/speaker.md.tpl -d output/speakers/ --log-level $(LOG_LEVEL)

program: bin/papercallfmt  ## Clear and re-generate program files based on accepted submissions (set JSON_FILE / LOG_LEVEL)
	-rm output/program/*.md
	bin/papercallfmt -s $(JSON_FILE) -t templates/program.md.tpl -d output/program/ --log-level $(LOG_LEVEL)

build: bin/papercallfmt   ## Build and print help output
	bin/papercallfmt --help

clean:   ## Clean generated files
	-rm output/speakers/*.md
	-rm output/program/*.md
