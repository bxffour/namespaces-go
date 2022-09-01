.PHONY: build
build:
	@go build -o ./bin/ns .
	@echo 'build succeeded'

.PHONY: assets/setup
assets/setup:
	@echo "creating rootfs' folder"
	mkdir -p /tmp/ns-process/rootfs
	@echo 'extracting to rootfs'
	tar -C /tmp/ns-process/rootfs -xf ./assets/rootfs/busybox.tar
