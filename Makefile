run:
	go run cmd/app/main.go
.PHONY: run

build:
	CGO_ENABLED=0 go build -o ./dist/app ./cmd/app
.PHONY: build

test:
	go test ./...
.PHONY: test

lint:
	@./scripts/linter.sh

install-git-tools:
	pip install pre-commit || yes
	rm -f .git/hooks/pre-commit
	rm -f .git/hooks/commit-msg
	ln -sr ./scripts/pre-commit .git/hooks
	ln -sr ./scripts/commit-msg .git/hooks
.PHONY: install-git-tools

swag:
	swag init -g cmd/app/main.go
	git add docs
.PHONY: swag


skip-checks:
	@touch skip.checks.flag
.PHONY: skip-checks


