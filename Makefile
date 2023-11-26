include ./common.mk

SERVICE_NAME = catCafe

.PHONY: build
build:
	sh build.sh

.PHONY: new
new:
	hz new \
	-module $(MODULE) \

.PHONY: gen
gen:

.PHONY: server
server:
	make build
	sed -i 's/\r//' ./output/bootstrap.sh
	cd output && sh bootstrap.sh

.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print