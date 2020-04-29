#!/usr/bin/env bash
.PHONY : fmt build

fmt:
	gofmt -l -w ./

build:
	go test .
