FROM golang:1.17

ENV GOLANG_CI_LINT_VERSION=v1.30.0

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

EXPOSE 80

WORKDIR /app