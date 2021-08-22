# GO related
GO      ?= go
LDFLAGS = -X "main.CommitHash=$(COMMIT_HASH)" -X "main.Tag=$(TAG)"

# re-build the image
up:
	$(GO) mod download
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags '$(LDFLAGS)' -o ./bin/main cmd/cli/main.go
	@docker-compose up -d --build

# remove the container
down:
	@docker-compose down

# execute binary from container
run:
	@docker-compose run --rm app

# Run tests
tests:
	@$(GO) test -race -v ./... -coverprofile=testdata/coverage.out

# View test coverage in browser
coverage:
	@$(GO) tool cover -html=testdata/coverage.out
