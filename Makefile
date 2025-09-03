# Makefile for Go project

# Detect the operating system and architecture.

include makefiles/osdetect.mk

# -----------------------------------------------------------------------------
# Variables
# -----------------------------------------------------------------------------

# "Simple expanded" variables (':=')

# PROGRAM_NAME is the name of the GIT repository.
PROGRAM_NAME := $(shell basename `git rev-parse --show-toplevel`)
MAKEFILE_PATH := $(abspath $(firstword $(MAKEFILE_LIST)))
MAKEFILE_DIRECTORY := $(shell dirname $(MAKEFILE_PATH))
TARGET_DIRECTORY := $(MAKEFILE_DIRECTORY)/target
DIST_DIRECTORY := $(MAKEFILE_DIRECTORY)/dist
BUILD_TAG := $(shell git describe --always --tags --abbrev=0  | sed 's/v//')
BUILD_ITERATION := $(shell git log $(BUILD_TAG)..HEAD --oneline | wc -l | sed 's/^ *//')
BUILD_VERSION := $(shell git describe --always --tags --abbrev=0 --dirty  | sed 's/v//')
GIT_REMOTE_URL := $(shell git config --get remote.origin.url)
GIT_REPOSITORY_NAME := $(shell basename `git rev-parse --show-toplevel`)
GIT_VERSION := $(shell git describe --always --tags --long --dirty | sed -e 's/\-0//' -e 's/\-g.......//')
GO_PACKAGE_NAME := $(shell echo $(GIT_REMOTE_URL) | sed -e 's|^git@github.com:|github.com/|' -e 's|\.git$$||' -e 's|Senzing|senzing|')
PATH := $(MAKEFILE_DIRECTORY)/bin:$(PATH)

# Recursive assignment ('=')

GO_OSARCH = $(subst /, ,$@)
GO_OS = $(word 1, $(GO_OSARCH))
GO_ARCH = $(word 2, $(GO_OSARCH))

# Conditional assignment. ('?=')
# Can be overridden with "export"
# Example: "export LD_LIBRARY_PATH=/path/to/my/senzing/er/lib"

GOBIN ?= $(shell go env GOPATH)/bin

# Tricky code.
# Accept a LD_LIBRARY_PATH environment variable, default to /opt/senzing/er/lib
# Then append path to oracle SDK

LD_LIBRARY_PATH ?= /opt/senzing/er/lib
LD_LIBRARY_PATH := $(LD_LIBRARY_PATH):/opt/oracle/instantclient_23_5

# Export environment variables.

.EXPORT_ALL_VARIABLES:

# -----------------------------------------------------------------------------
# The first "make" target runs as default.
# -----------------------------------------------------------------------------

.PHONY: default
default: help

# -----------------------------------------------------------------------------
# Operating System / Architecture targets
# -----------------------------------------------------------------------------

-include makefiles/$(OSTYPE).mk
-include makefiles/$(OSTYPE)_$(OSARCH).mk


.PHONY: hello-world
hello-world: hello-world-osarch-specific

# -----------------------------------------------------------------------------
# Dependency management
# -----------------------------------------------------------------------------

.PHONY: dependencies-for-development
dependencies-for-development: dependencies-for-development-osarch-specific
	@go install github.com/bombsimon/wsl/v5/cmd/wsl@latest
	@go install github.com/daixiang0/gci@latest
	@go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest
	@go install github.com/vladopajic/go-test-coverage/v2@latest
	@go install golang.org/x/tools/cmd/godoc@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@go install mvdan.cc/gofumpt@latest
	@docker-compose -f docker-compose.test.yaml pull 2>/dev/null || true
	@sudo npm install -g cspell@latest


.PHONY: dependencies
dependencies:
	@go get -u ./...
	@go get -t -u ./...
	@go mod tidy

# -----------------------------------------------------------------------------
# Setup
# -----------------------------------------------------------------------------

.PHONY: setup
setup: setup-osarch-specific

# -----------------------------------------------------------------------------
# Lint
# -----------------------------------------------------------------------------

.PHONY: lint
lint: golangci-lint govulncheck cspell

# -----------------------------------------------------------------------------
# Build
# -----------------------------------------------------------------------------

PLATFORMS := darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64
$(PLATFORMS):
	$(info Building $(TARGET_DIRECTORY)/$(GO_OS)-$(GO_ARCH)/$(PROGRAM_NAME))
	@GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go build -o $(TARGET_DIRECTORY)/$(GO_OS)-$(GO_ARCH)/$(PROGRAM_NAME)


.PHONY: build
build: build-osarch-specific


.PHONY: build-all $(PLATFORMS)
build-all: $(PLATFORMS)
	@mv $(TARGET_DIRECTORY)/windows-amd64/$(PROGRAM_NAME) $(TARGET_DIRECTORY)/windows-amd64/$(PROGRAM_NAME).exe
	@mv $(TARGET_DIRECTORY)/windows-arm64/$(PROGRAM_NAME) $(TARGET_DIRECTORY)/windows-arm64/$(PROGRAM_NAME).exe

# -----------------------------------------------------------------------------
# Run
# -----------------------------------------------------------------------------

.PHONY: run
run: run-osarch-specific

# -----------------------------------------------------------------------------
# Test
# -----------------------------------------------------------------------------

.PHONY: test
test: test-osarch-specific

# -----------------------------------------------------------------------------
# Coverage
# -----------------------------------------------------------------------------

.PHONY: coverage
coverage: coverage-osarch-specific


.PHONY: check-coverage
check-coverage: export SENZING_LOG_LEVEL=TRACE
check-coverage:
	@go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	@${GOBIN}/go-test-coverage --config=.github/coverage/testcoverage.yaml

# -----------------------------------------------------------------------------
# Documentation
# -----------------------------------------------------------------------------

.PHONY: documentation
documentation: documentation-osarch-specific

# -----------------------------------------------------------------------------
# Clean
# -----------------------------------------------------------------------------

.PHONY: clean
clean: clean-osarch-specific
	@go clean -cache
	@go clean -testcache

# -----------------------------------------------------------------------------
# Utility targets
# -----------------------------------------------------------------------------

.PHONY: help
help:
	$(info Build $(PROGRAM_NAME) version $(BUILD_VERSION)-$(BUILD_ITERATION))
	$(info Makefile targets:)
	@$(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs


.PHONY: print-make-variables
print-make-variables:
	@$(foreach V,$(sort $(.VARIABLES)), \
		$(if $(filter-out environment% default automatic, \
		$(origin $V)),$(info $V=$($V) ($(value $V)))))


.PHONY: update-pkg-cache
update-pkg-cache:
	@GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get $(GO_PACKAGE_NAME)@$(BUILD_TAG)

# -----------------------------------------------------------------------------
# Specific programs
# -----------------------------------------------------------------------------

.PHONY: bearer
bearer:
	@bearer scan --config-file .github/linters/bearer.yml .


.PHONY: cspell
cspell:
	@cspell lint --dot .


.PHONY: exhaustruct
exhaustruct:
	exhaustruct ./...


.PHONY: gofumpt
gofumpt:
	gofumpt -d ./**/*.go


.PHONY: golangci-lint
golangci-lint:
	@${GOBIN}/golangci-lint run --config=.github/linters/.golangci.yaml


.PHONY: govulncheck
govulncheck:
	@${GOBIN}/govulncheck ./...

# -----------------------------------------------------------------------------
# Fixers
# -----------------------------------------------------------------------------

.PHONY: fix
fix: fix-asciicheck
fix: fix-bidichk
fix: fix-canonicalheader
# fix: fix-copyloopvar
fix: fix-cyclop
fix: fix-dupword
# fix: fix-durationcheck
# fix: fix-err113
fix: fix-errchkjson
fix: fix-errname
fix: fix-errorlint
# fix: fix-exhaustive
# fix: fix-exhaustruct
# fix: fix-exptostd
# fix: fix-fatcontext
# fix: fix-ginkgolinter
# fix: fix-gocheckcompilerdirectives
# fix: fix-gochecknoglobals
# fix: fix-godot
fix: fix-gofumpt
# fix: fix-grouper
# fix: fix-ifacecheck
# fix: fix-interfacebloat
fix: fix-inamedparam
# fix: fix-ireturn
fix: fix-loggercheck
# fix: fix-maintidx
# fix: fix-mirror
# fix: fix-mnd
fix: fix-nakedret
# fix: fix-nilerr
fix: fix-nilnesserr
fix: fix-nilnil
fix: fix-paralleltest
fix: fix-perfsprint
# fix: fix-predeclared
# fix: fix-protogetter
# fix: fix-rowserrcheck
fix: fix-tagalign
fix: fix-tagliatelle
# fix: fix-testableexamples
fix: fix-testifylint
# fix: fix-testpackage
# fix: fix-thelper
# fix: fix-usestdlibvars
fix: fix-usetesting
# fix: fix-whitespace
fix: fix-wrapcheck
fix: fix-wsl
	$(info fixes complete)


.PHONY: fix-asciicheck
fix-asciicheck:
	@asciicheck --fix ./...


.PHONY: fix-bidichk
fix-bidichk:
	@bidichk --fix ./...


.PHONY: fix-canonicalheader
fix-canonicalheader:
	@canonicalheader --fix ./...


.PHONY: fix-copyloopvar
fix-copyloopvar:
	@copyloopvar --fix ./...


.PHONY: fix-cyclop
fix-cyclop:
	@cyclop --fix ./...


.PHONY: fix-dupword
fix-dupword:
	@dupword --fix ./...


.PHONY: fix-durationcheck
fix-durationcheck:
	@durationcheck --fix ./...


.PHONY: fix-err113
fix-err113:
	@err113 --fix ./...


.PHONY: fix-errchkjson
fix-errchkjson:
	@errchkjson --fix ./...


.PHONY: fix-errname
fix-errname:
	@errname --fix ./...


.PHONY: fix-errorlint
fix-errorlint:
	@go-errorlint --fix ./...


.PHONY: fix-exhaustive
fix-exhaustive:
	@go-exhaustive --fix ./...


.PHONY: fix-exhaustruct
fix-exhaustruct:
	@go-exhaustruct --fix ./...


.PHONY: fix-exptostd
fix-exptostd:
	@go-exptostd --fix ./...


.PHONY: fix-fatcontext
fix-fatcontext:
	@go-fatcontext -fix ./...


.PHONY: fix-ginkgolinter
fix-ginkgolinter:
	@go-ginkgolinter --fix ./...


.PHONY: fix-gocheckcompilerdirectives
fix-gocheckcompilerdirectives:
	@go-gocheckcompilerdirectives --fix ./...


.PHONY: fix-gochecknoglobals
fix-gochecknoglobals:
	@go-gochecknoglobals --fix ./...


.PHONY: fix-godot
fix-godot:
	@go-godot --fix ./...


.PHONY: fix-gofumpt
fix-gofumpt:
	@gofumpt -w ./**/*.go


.PHONY: fix-grouper
fix-grouper:
	@grouper --fix ./...


.PHONY: fix-ifacecheck
fix-ifacecheck:
	@ifacecheck --fix ./...


.PHONY: fix-interfacebloat
fix-interfacebloat:
	@interfacebloat --fix ./...


.PHONY: fix-inamedparam
fix-inamedparam:
	@inamedparam --fix ./...


.PHONY: fix-ireturn
fix-ireturn:
	@ireturn --fix ./...


.PHONY: fix-loggercheck
fix-loggercheck:
	@loggercheck --fix ./...


.PHONY: fix-maintidx
fix-maintidx:
	@maintidx --fix ./...


.PHONY: fix-mirror
fix-mirror:
	@mirror --fix ./...


.PHONY: fix-mnd
fix-mnd:
	@mnd --fix ./...


.PHONY: fix-nakedret
fix-nakedret:
	@nakedret --fix ./...


.PHONY: fix-nilerr
fix-nilerr:
	@nilerr --fix ./...


.PHONY: fix-nilnesserr
fix-nilnesserr:
	@nilnesserr --fix ./...


.PHONY: fix-nilnil
fix-nilnil:
	@nilnil --fix ./...


.PHONY: fix-paralleltest
fix-paralleltest:
	@paralleltest --fix ./...


.PHONY: fix-perfsprint
fix-perfsprint:
	@perfsprint --fix ./...


.PHONY: fix-predeclared
fix-predeclared:
	@predeclared --fix ./...


.PHONY: fix-protogetter
fix-protogetter:
	@protogetter --fix ./...


.PHONY: fix-rowserrcheck
fix-rowserrcheck:
	@rowserrcheck --fix ./...


.PHONY: fix-tagalign
fix-tagalign:
	@tagalign --fix ./...


.PHONY: fix-tagliatelle
fix-tagliatelle:
	@tagliatelle --fix ./...


.PHONY: fix-testableexamples
fix-testableexamples:
	@testableexamples --fix ./...


.PHONY: fix-testifylint
fix-testifylint:
	@testifylint --fix ./...


.PHONY: fix-testpackage
fix-testpackage:
	@testpackage --fix ./...


.PHONY: fix-thelper
fix-thelper:
	@thelper --fix ./...


.PHONY: fix-usestdlibvars
fix-usestdlibvars:
	@usestdlibvars --fix ./...


.PHONY: fix-usetesting
fix-usetesting:
	@usetesting --fix ./...


.PHONY: fix-whitespace
fix-whitespace:
	@whitespace --fix ./...


.PHONY: fix-wrapcheck
fix-wrapcheck:
	@wrapcheck --fix ./...


.PHONY: fix-wsl
fix-wsl:
	@wsl --fix ./...
