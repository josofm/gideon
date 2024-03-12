pjname=gideon
version?=latest
img=$(pjname):$(version)
rundev=docker run --rm $(img)
runbuild=docker run --rm -e CGO_ENABLED=0 GOOS=linux -e GOARCH=amd64 $(img)
rundevti=docker run -ti --rm $(img)
runcompose=docker compose run --rm


.PHONY: all
all: check build
guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

.PHONY: imagedev
imagedev:
	docker build --target devimage . -t $(img)

.PHONY: image
image:
	docker build --target prodimage --build-arg version=$(version) -t $(img)

.PHONY: check
check: imagedev
	$(rundev) ./hack/check.sh $(testfile) $(testname)

.PHONY: shell
shell: imagedev
	$(rundevti) sh



