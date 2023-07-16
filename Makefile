# ---------------------------------------- COMMON START ----------------------------------------------------------------
.PHONY: test
test:
	CGO_ENABLED=1 go test -v -race -count=1 ./...

.PHONY: lint
lint:
	@golangci-lint run
# ---------------------------------------- COMMON END ------------------------------------------------------------------

# ---------------------------------------- PROTOBUF START --------------------------------------------------------------
NAMES=auth

.PHONY: all
all: clean format lint gen

.PHONY: gen
gen:
	@$(GOPATH)/bin/buf generate
	@for name in $(NAMES); do \
  		cd $(CURDIR)/gen/go/$$name && \
  		go mod init github.com/theartofedu/grpc-contracts-repo/gen/go/$$name && go mod tidy; \
  	done

.PHONY: tags
tags:
	@for name in $(NAMES); do \
  		version=$$(cat $(CURDIR)/proto/$$name/v) && echo "work with $$name and $$version" && \
  		tag=gen/go/$$name/v$$version && echo "tag: $$tag" && \
  		if [[ ! $$(git tag -l "$$tag") ]]; then \
  		  	git tag -a "$$tag" -m "" && \
  		  	git push origin "$$tag" -o ci.skip; \
		fi; \
	done

.PHONY: lint
lint:
	@$(GOPATH)/bin/buf lint

.PHONY: format
format:
	@$(GOPATH)/bin/buf format

.PHONY: clean
clean:
	@rm -rf ./gen || true
# ---------------------------------------- PROTOBUF END ----------------------------------------------------------------