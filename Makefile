DIST_DIR = "public"

.PHONY: all
all: build

.PHONY: build
build:
	hugo

.PHONY: dev
dev:
	hugo server --buildDrafts

.PHONY: server
server:
	hugo server


.PHONY: clean
clean:
	rm -rf $(DIST_DIR)

.PHONY: re
re: clean build

.PHONY: gzip
gzip:
	gzip -k -9 -r -f $(DIST_DIR)
