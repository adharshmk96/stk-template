##########################
### Version Commands
##########################

patch:
	$(eval NEW_TAG := $(shell git semver patch --dryrun))
	$(call update_file)
	@git semver patch

minor:
	$(eval NEW_TAG := $(shell git semver minor --dryrun))
	$(call update_file)
	@git semver minor

major:
	$(eval NEW_TAG := $(shell git semver major --dryrun))
	$(call update_file)
	@git semver major

publish:
	@git push origin $(shell git semver get)


##########################
### Build Commands
##########################

BINARY_NAME=app

build:
	@go build -o ./out/$(BINARY_NAME) -v

run: 
	@go run . serve -p 8080

test:
	@go test ./... -coverprofile=coverage.out

coverage:
	@go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

testci:
	@go test ./... -coverprofile=coverage.out

clean:
	@go clean
	@rm -f ./out/$(BINARY_NAME)
	@rm -f coverage.out
	@rm -rf .keys
	@rm -f auth_database.db

deps:
	@go mod download

tidy:
	@go mod tidy

lint:
	@golangci-lint run --enable-all

vet:
	@go vet

clean-branch:
	@git branch | egrep -v "(^\*|main|master)" | xargs git branch -D

	
##########################
### Helpers
##########################

define update_file
    @echo "updating files to version $(NEW_TAG)"
    @sed -i.bak "s/var version = \"[^\"]*\"/var version = \"$(NEW_TAG)\"/g" ./cmd/root.go
    @rm cmd/root.go.bak
    @git add cmd/root.go
    @git commit -m "bump version to $(NEW_TAG)" > /dev/null
endef

##########################
### Setup Commands
##########################

init: deps keygen initgithooks mockgen
	@echo "Project initialized."

initci: deps keygen
	@echo "Project initialized for CI."

initgithooks:
	@git config core.hooksPath .githooks

mockgen:
	@rm -rf ./mocks
	@mockery --all	

