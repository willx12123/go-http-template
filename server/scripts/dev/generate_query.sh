#!/bin/bash

GENERATE_DIR="./cmd/generate"

cd $GENERATE_DIR || exit

echo "Start Generating GORM......"
go run .
