GOBIN ?= $$PWD/bin

define build
    mkdir -p $(GOBIN)
    GOGC=off GOBIN=$(GOBIN) \
		go install -v \
		-mod=vendor \
		-gcflags='-e' \
		$(1)
endef

build-x0:
	$(call build, ./cmd/x0)

build-x1event:
	$(call build, ./cmd/x1event)

build-j0:
	$(call build, ./cmd/j0)