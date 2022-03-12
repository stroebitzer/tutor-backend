IMAGE_REPOSITORY = ueber
APPLICATION_NAME = tutor-backend
BUILD_VERSION = 0.0.1

.PHONY: test
test:
	go test ./... 

.PHONY: build
build: test
# TODO do this CGO_ENABLED everywhere, also in Dockerfiles, and ensure ubuntu is even more happy
	CGO_ENABLED=0 GOOS=linux go build -o ${APPLICATION_NAME}

.PHONY: run
run: build
# TODO
# TRAINING_DIR=/home/hubert/git/tutor-training TRAINING_FILE=.training.yaml ./${APPLICATION_NAME}
	./${APPLICATION_NAME}

# TODO
# .PHONY: release
# release: build
# 	git tag ${BUILD_VERSION}
# 	git push --tags

.PHONY: docker-build
docker-build: build
	docker build -t ${IMAGE_REPOSITORY}/${APPLICATION_NAME}:${BUILD_VERSION} .

.PHONY: docker-run
docker-run: docker-build
	docker run -it --rm -p 8080:8080 --name ${APPLICATION_NAME} ${IMAGE_REPOSITORY}/${APPLICATION_NAME}:${BUILD_VERSION}

.PHONY: docker-push
docker-push: docker-build
	docker push ${IMAGE_REPOSITORY}/${APPLICATION_NAME}:${BUILD_VERSION}

.PHONY: k8s-patch
k8s-patch: docker-push
	kubectl -n tutor rollout restart deployment tutor  

