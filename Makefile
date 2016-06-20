.PHONY=docker

docker:
	docker run --rm \
	  -v "$$(pwd):/src" \
	  -v /var/run/docker.sock:/var/run/docker.sock \
	  -e COMPRESS_BINARY=true \
	  centurylink/golang-builder \
	  $${BUILDKITE_DOCKER_COMPOSE_BUILD_TAG}
