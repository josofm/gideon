FROM golang:1.12-alpine
LABEL mainterner 'josofm'


COPY go.mod go.sum ./

COPY . .