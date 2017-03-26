help:
	@echo 'Makefile for go-search-engine                                     '
	@echo '                                                                      '
	@echo 'Usage:                                                                '
	@echo '   test                                Run project tests              '

test:
	@go test ./...

.PHONY: help test
