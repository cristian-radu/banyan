BINARY_NAME=banyan
IMAGE_REGISTRY=cristianradu
ABBREVIATED_COMMIT_HASH=$(shell git log -1 --format=%h)

artifact: generate-code docker-build docker-push 
generate-code:
	bash $(CURDIR)/build/code-gen.sh
docker-build:
	docker build . -t $(IMAGE_REGISTRY)/$(BINARY_NAME):$(ABBREVIATED_COMMIT_HASH)
docker-push:
	docker push $(IMAGE_REGISTRY)/$(BINARY_NAME):$(ABBREVIATED_COMMIT_HASH)
