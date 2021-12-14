GOCMD				=go
TEST				?=$$(go list ./... |grep -v 'vendor')
INSTALL_DIR			=~/.terraform.d/plugins
BINARY				=terraform-provider-equinix
SWEEP_DIR			?=./equinix
SWEEPARGS			=timeout 60m
ACCTEST_TIMEOUT		?= 180m
ACCTEST_PARALLELISM ?= 20
ACCTEST_COUNT       ?= 1

ifneq ($(origin TESTS_REGEXP), undefined)
	TESTARGS = -run='$(TESTS_REGEXP)'
endif

default: clean build test

all: default
	
test:
	echo $(TEST) | \
		xargs -t ${GOCMD} test -v $(TESTARGS) -timeout=10m

testacc:
	TF_ACC=1 TF_SCHEMA_PANIC_ON_ERROR=1 ${GOCMD} test $(TEST) -v -count $(ACCTEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(TESTARGS) -timeout $(ACCTEST_TIMEOUT)

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test $(SWEEP_DIR) -v -sweep=all $(SWEEPARGS)

build:
	${GOCMD} build -o ${BINARY}

install: test build
	@if [ -d ${INSTALL_DIR} ]; then \
		echo "==> [INFO] installing in ${INSTALL_DIR} directory"; \
		cp ${BINARY} ${INSTALL_DIR}; \
	else \
		echo "==> [ERROR] installation plugin directory ${INSTALL_DIR} does not exist"; \
	fi

clean:
	${GOCMD} clean
	rm -f ${BINARY}

.PHONY: test testacc build install clean
