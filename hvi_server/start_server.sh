#!/bin/bash
cd src
if [[  "$1" == "test" ]]; then
	mkdir -p ../coverage-report
	go test -v -coverprofile=../coverage-report/coverage.out ./...
	go tool cover -html=../coverage-report/coverage.out -o ../coverage-report/coverage.html
else
        set -a
        source ../../.env
        set +a
        go run main.go
fi
cd ../
