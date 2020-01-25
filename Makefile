PROJECT_DIR   = $(shell readlink -f .)
BUILD_DIR     = $(PROJECT_DIR)/build
SERVER_DIR    = $(PROJECT_DIR)/cmd/kahle-server
SERVER_BIN    = $(BUILD_DIR)/kahle-server

GO           ?= go
RICHGO       ?= rich$(GO)

VERSION       = $(shell git describe --tags --always --dirty)
LD_FLAGS      = "-X github.com/wavesoftware/go-kahle/internal/base.Version=$(VERSION)"

.PHONY: default
default: binary

.PHONY: builddeps
builddeps:
	@GO111MODULE=off $(GO) get github.com/kyoh86/richgo
	@GO111MODULE=off $(GO) get github.com/mgechev/revive

.PHONY: builddir
builddir:
	@mkdir -p build

.PHONY: clean
clean: builddeps
	@echo "🛁 Cleaning"
	@rm -frv $(BUILD_DIR)

.PHONY: check
check: builddeps
	@echo "🛂 Checking"
	revive -config revive.toml -formatter stylish ./...

.PHONY: test
test: builddir check
	@echo "✔️ Testing"
	$(RICHGO) test -v -covermode=count -coverprofile=build/coverage.out ./...

.PHONY: binary
binary: test
	@echo "🔨 Building forwarder"
	env CGO_ENABLED=0 GOOS=linux $(RICHGO) build \
		-ldflags $(LD_FLAGS) \
		-o $(SERVER_BIN) \
		$(SERVER_DIR)
