#!/usr/bin/env make

.PHONY: test

test:
	ginkgo generics util config
