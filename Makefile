GOBIN ?= $$PWD/bin

define test-lib
	GOGC=off go test $(TEST_FLAGS) $$(go list ./... | grep -v -e /vendor/ -e /cmd/) | grep -v '\[no test files\]'
endef

define build
    mkdir -p $(GOBIN)
    GOGC=off GOBIN=$(GOBIN) \
		go install -v \
		-mod=vendor \
		-gcflags='-e' \
		$(1)
endef

build-j0:
	$(call build, ./cmd/j0)

test-lib:
	$(call test-lib)