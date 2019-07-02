BINARY_NAME=banyan

API_VERSION=v1alpha1
APIS_PATH=github.com/cristian-radu/$(BINARY_NAME)/pkg/apis

CODE_GEN_DIR=$(shell grep k8s.io/code-generator go.mod | sed -E -e 's/^[[:blank:]]|[[:blank:]]$$//' -e s'/[[:blank:]]/@/')
CODE_GEN_PATH=${GOPATH}/pkg/mod/$(CODE_GEN_DIR)

GIT_COMMIT=$(shell git log -1 --format=%h)

IMAGE_REGISTRY=cristianradu

code-gen-download:
	go mod download k8s.io/code-generator

code-gen-run: code-gen-download
	chmod +x $(CODE_GEN_PATH)/generate-groups.sh
	$(CODE_GEN_PATH)/generate-groups.sh all ./client $(APIS_PATH) $(BINARY_NAME):$(API_VERSION) -o ./
	mv $(CURDIR)/$(APIS_PATH)/$(BINARY_NAME)/$(API_VERSION)/zz_generated.deepcopy.go $(CURDIR)/pkg/apis/$(BINARY_NAME)/$(API_VERSION)
	rm -rf $(CURDIR)/github.com

docker-build:
	docker build . -t $(IMAGE_REGISTRY)/$(BINARY_NAME):$(GIT_COMMIT)

docker-push: docker-build
	docker push $(IMAGE_REGISTRY)/$(BINARY_NAME):$(GIT_COMMIT)