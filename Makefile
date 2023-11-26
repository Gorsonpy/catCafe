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
	hz update -idl ./idl/api.thrift
	hz update -idl ./idl/multiLedger.thrift
	hz update -idl ./idl/ledger.thrift
	hz update -idl ./idl/consumption.thrift
	hz update -idl ./idl/asr.thrift

.PHONY: server
server:
	make build
	sed -i 's/\r//' ./output/bootstrap.sh
	cd output && sh bootstrap.sh

.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print