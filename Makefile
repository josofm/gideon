SERVICE_NAME := gideon
version?=latest
img=gideon
uid=$(shell id -u $$USER)
gid=$(shell id -g $$USER)
wd=$(shell pwd)
appvol=$(wd):/app
modcachedir=$(wd)/.gomodcachedir
cachevol=$(modcachedir):/go/pkg/mod
rundev=docker run --rm -v $(appvol) -v $(cachevol) $(img)
runbuild=docker run --rm -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 -v $(appvol) -v $(cachevol) $(SERVICE_NAME)
SOURCE_DIR=/home/joso/workspace/projects/gideon
RUN_GO=--rm -e LOGGER_LEVEL=$(log) -v `pwd`:$(SOURCE_DIR) -v $(cachevol) -w $(SOURCE_DIR) $(SERVICE_NAME)
runcompose=docker-compose run --rm -v $(appvol) -v $(cachevol)
cov=coverage.out
covhtml=coverage.html

all: check build
guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi
# WHY: If cache dir does not exist it is mapped inside container as root
# If it exists it is mapped belonging to the non-root user inside the container

modcache:
	@mkdir -p $(modcachedir)

imagedev:
	docker build . -t $(SERVICE_NAME) -f ./hack/Dockerfile

build: modcache imagedev
	$(runbuild) go build -v -ldflags "-w -s -X main.Version=$(version)" -o ./cmd/gideon/gideon ./cmd/gideon

check: modcache imagedev
	docker run $(RUN_GO) ./hack/check.sh $(suite) $(test)

start-compose:
	docker-compose pull --ignore-pull-failures
	@ if [ "${ci}" = "true" ]; then \
	    docker-compose up -d; \
	else \
		docker-compose -f docker-compose.yml up -d; \
	fi
check-integration: build start-compose
	$(runcompose) --entrypoint "./hack/check-integration.sh $(workdir)/$(test)" gideon
	@ if [ "${log}" != "debug" ]; then \
		docker-compose kill; \
		docker-compose rm -fv; \
	fi
stop:
	docker-compose kill
	docker-compose rm -fv
coverage: modcache check
	$(rundev) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html
static-analysis: modcache imagedev
	$(rundev) golangci-lint run ./...
modtidy: modcache imagedev
	$(rundev) go mod tidy
fmt: modcache imagedev
	$(rundev) gofmt -w -s -l .
githooks:
	@echo "copying git hooks"
	@mkdir -p .git/hooks
	@cp hack/githooks/pre-commit .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "git hooks copied"
shell: modcache imagedev
	$(rundev) sh
run: imagedev
	$(runcompose) --service-ports --entrypoint "go run ./cmd/gideon/gideon.go" gideon


