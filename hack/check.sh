#!/bin/bash


if [ $# -eq 0 ]
then
    go test -tags unit -race -covermode=atomic -timeout 30s ./...
    exit
fi

if [ $# -eq 1 ]
    then
        go test -tags unit -run $1 -race -covermode=atomic -timeout 30s ./...
        exit
fi

go test -tags unit -run $1 -m $2 -race -covermode=atomic -timeout 30s ./...