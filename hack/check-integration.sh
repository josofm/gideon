#!/bin/bash

sleep 5

if [ $# -eq 0 ]
then
    go test -tags integration -p 1 -covermode=atomic -timeout 50s ./...
    exit
fi
if [ $# -eq 1 ]
then
    go test -tags integration -p 1 -run $1 -covermode=atomic -timeout 50s ./...
    exit
fi
go test -tags integration -p 1 -run $1 -testify.m $2 -covermode=atomic -timeout 50s ./...