IMAGE = harbor.home.twofei.com/sgits:latest

.PHONY: build-linux
build-linux:
	GOOS=linux go build
.PHONY: build-image
build-image: build-linux
	docker build -t ${IMAGE} .
.PHONY: push-image
push-image: build-image
	docker push ${IMAGE}
