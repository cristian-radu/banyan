BINARY_NAME=banyan
IMAGE_REGISTRY=cristianradu
LAST_GIT_COMMIT_ID=$(shell git log -1 --format=%h)

generate-code:
	bash $(CURDIR)/build/code-gen.sh

docker-build: generate-code
	docker build . -t $(IMAGE_REGISTRY)/$(BINARY_NAME):$(LAST_GIT_COMMIT_ID)

docker-push: docker-build
	docker push $(IMAGE_REGISTRY)/$(BINARY_NAME):$(LAST_GIT_COMMIT_ID)