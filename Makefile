OUTPUT?=build/
CGO_ENABLED ?= 0

all: build install

.PHONY: all

# The below include contains the tools.
include tools.mk

tools:=addr bbm bcparser bcscan bcw genesis methodid orgid relay sigorg smccheck smcpack txparse txsampleV2 walv1tov2
build: $(tools) bcc

.PHONY: build

$(tools):
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(OUTPUT)$@ ./$@
.PHONY: $(tools)

bcc:
	CGO_ENABLED=$(CGO_ENABLED) go build -o $(OUTPUT) ./bcc/cmd/...
.PHONY: bcc

install:
	cp -rf ./bundle/.config ./build/
.PHONY: install

dist:
	@sh -c "'$(CURDIR)/scripts/dist.sh'"
.PHONY: dist