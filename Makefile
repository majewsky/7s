PKG    = github.com/majewsky/7s
PREFIX = /usr

GO            = GOPATH=$(CURDIR)/.gopath GOBIN=$(CURDIR)/build go
GO_BUILDFLAGS =
GO_LDFLAGS    = -s -w

all: bindata.go FORCE
	$(GO) install $(GO_BUILDFLAGS) -ldflags '$(GO_LDFLAGS)' '$(PKG)'
bindata.go: static/*
	go-bindata static/*

install: FORCE all
	install -D -m 0755 build/7s "$(DESTDIR)$(PREFIX)/bin/7s"

example-%: examples/%.7s | all
	./build/7s $< :8080


vendor: FORCE
	# vendoring by https://github.com/holocm/golangvend
	golangvend

.PHONY: FORCE
