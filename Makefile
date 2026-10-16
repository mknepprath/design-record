GO ?= go
RECORD = $(GO) run ./tools/record

.PHONY: hooks new index check test

hooks: ## once per clone: make check runs before every push
	git config core.hooksPath .githooks

new: ## lists the kinds; with KIND and TITLE, creates a document from its template
	@$(RECORD) new "$(KIND)" "$(TITLE)" "$(PROJECT)" "$(SOURCE)"

index: ## regenerate INDEX.md
	@$(RECORD) index

check: ## the rules, checked
	@test -z "$$(gofmt -l tools)" || (gofmt -l tools; echo "gofmt needed"; exit 1)
	@$(GO) vet ./tools/...
	@$(GO) test ./tools/...
	@$(RECORD) check

test:
	$(GO) test ./tools/...
