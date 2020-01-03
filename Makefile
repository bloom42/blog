.PHONY: all
all: build

.PHONY: build
build:
	hugo

.PHONY: dev
dev:
	hugo server --buildDrafts
