FROM golang:1.21 as devimage

ENV GOLANG_CI_LINT_VERSION=V1.56.2
ENV GIT_TERMINAL_PROMPT=1
ENV GOPROXY=direct

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

EXPOSE 80
WORKDIR /app
COPY go.mod /app
RUN go mod download
RUN go mod tidy
COPY . /app

FROM devimage as buildimage
ARG GOLANG_CI_LINT_VERSION
COPY --from=devimage /app ./
ENV CGO_ENABLED=0
RUN go build -a  -installsufix cgo -ldflags "-w -s  -X main.Version=$version" -o .cmd/gideon

FROM alpine:3.9 as prodimage
RUN apk --no-cache update && \
	apk --no-cache add ca-certificates tzdata && \
	rm -rf /var/cache/apk/*
WORKDIR /

COPY --from=buildimage /app/cmd/gideon/gideon /app/gideon

EXPOSE 80

ENTRYPOINT ["/app/gideon"]