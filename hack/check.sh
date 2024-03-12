#!/bin/bash

set -o errexit
set -o nounset

run_cmd="go test -timeout 30s -tags=unit -p 1"

if [ $# -eq 0 ]; then
    $run_cmd ./...
    exit
fi

if [ $# -eq 1 ]; then
    testfile=$1
    echo "Running test $testfile"
    $run_cmd $testfile
fi

testfile=$1
testname=$2
echo "Running testname $testname on testfile $testfile"
$run_cmd $testfile -run $testname