#!/bin/bash
set -a
source ../.env
set +a
cd src
go run main.go
cd ../
