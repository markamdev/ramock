# Binaries
GO = go

# Params
BUILD_DIR = $(shell pwd)/build

# Go compiler should always check if it's necessary to re-buld binary
.PHONY: ramock docker

# Default target
all: ramock config

$(BUILD_DIR):
	@echo -- BUILD DIR --
	@mkdir -p $(BUILD_DIR)

ramock: $(BUILD_DIR)
	@echo -- ramock --
	@cd cmd/ramock && $(GO) build -o $(BUILD_DIR)/ ./

config: $(BUILD_DIR)
	@echo -- CONFIG --
	@cp ./data/endpoins-example.yaml $(BUILD_DIR)/endpoints.yaml

clean:
	@echo -- CLEAN --
	@rm -rf $(BUILD_DIR)

docker:
	@echo -- DOCKER --
	@docker build -t markamdev/ramock -f Dockerfile .

podman:
	@echo -- PODMAN --
	@podman build -t markamdev/ramock -f Dockerfile .

# Targets for publishing images to Docker Hub
# docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
# docker buildx rm builder
# docker buildx create --name builder --driver docker-container --use
# docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t markamdev/ramock:latest -t markamdev/ramock:0.5 --push -f Dockerfile .

.publish_ramock:
	$(eval G_VER := $(shell cat data/ramock.VERSION | head -n 1))
	@echo -- DOCKERHUB PUBLISHING : ramock v $(G_VER) --
	@echo "INFO: multiarch build requires multiarch/qemu-user-static"
	@echo "Run it using: docker run --rm --privileged multiarch/qemu-user-static --reset -p yes"
	@docker buildx create --name ramockbuilder --driver docker-container --use
	@docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 \
		-t markamdev/ramock:latest -t markamdev/ramock:$(G_VER) --push -f Dockerfile .
	@docker buildx rm ramockbuilder
