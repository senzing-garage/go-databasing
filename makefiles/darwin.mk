# Makefile extensions for darwin.

# -----------------------------------------------------------------------------
# Variables
# -----------------------------------------------------------------------------

SENZING_TOOLS_DATABASE_URL ?= sqlite3://na:na@/tmp/sqlite/G2C.db

# -----------------------------------------------------------------------------
# OS specific targets
# -----------------------------------------------------------------------------

.PHONY: build-osarch-specific
build-osarch-specific: darwin/amd64


.PHONY: clean-osarch-specific
clean-osarch-specific:
	@docker rm --force $(DOCKER_CONTAINER_NAME) 2> /dev/null || true
	@docker rmi --force $(DOCKER_IMAGE_NAME) $(DOCKER_BUILD_IMAGE_NAME) 2> /dev/null || true
	@rm -rf $(TARGET_DIRECTORY) || true
	@rm -f $(GOPATH)/bin/$(PROGRAM_NAME) || true


.PHONY: hello-world-osarch-specific
hello-world-osarch-specific:
	@echo "Hello World, from darwin."


.PHONY: package-osarch-specific
package-osarch-specific:
	@echo No packaging for darwin.


.PHONY: run-osarch-specific
run-osarch-specific:
	@go run -exec macos_exec_dyld.sh main.go


.PHONY: setup-osarch-specific
setup-osarch-specific:
	@rm -rf /tmp/sqlite
	@mkdir  /tmp/sqlite
	@touch  /tmp/sqlite/G2C.db


.PHONY: test-osarch-specific
test-osarch-specific:
	@go test -v -p 1 ./...

# -----------------------------------------------------------------------------
# Makefile targets supported only by this platform.
# -----------------------------------------------------------------------------

.PHONY: only-darwin
only-darwin:
	@echo "Only darwin has this Makefile target."
