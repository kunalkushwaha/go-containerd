GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
BUILD_IMAGE := go-containerd_dev$(if $(GIT_BRANCH_CLEAN),:$(GIT_BRANCH_CLEAN))


dbuild: buildimage
	docker run --rm -it --privileged=true -v $(CURDIR):/go/src/github.com/kunalkushwaha/go-containerd -v ${PWD}/bin:/cbin $(BUILD_IMAGE) go build  -o /cbin/goctr

buildimage:
	docker build -t $(BUILD_IMAGE) .
