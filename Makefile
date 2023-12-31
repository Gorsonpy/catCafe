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
	hz update -idl ./idl/membership.thrift
	hz update -idl ./idl/cat.thrift
	hz update -idl ./idl/appointment.thrift
	hz update -idl ./idl/order.thrift
.PHONY: server
server:
	make build
	sed -i 's/\r//' ./output/bootstrap.sh
	cd output && sh bootstrap.sh

.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print