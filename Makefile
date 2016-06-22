.PHONY=clean

default: polyglot-co-demo-backend

polyglot-co-demo-backend:
	docker run --rm \
	  -v "$$(pwd):/src" \
	  -v /var/run/docker.sock:/var/run/docker.sock \
	  -e COMPRESS_BINARY=true \
	  centurylink/golang-builder

eb.zip: polyglot-co-demo-backend
	mkdir -p eb
	cp Dockerfile-production eb/Dockerfile
	cp -r templates static polyglot-co-demo-backend eb/
	echo '{"AWSEBDockerrunVersion": "1","Ports": [{"ContainerPort": "8080"}]}' > eb/Dockerrun.aws.json
	docker run --rm -v "$$(pwd):/src" alpine sh -c 'apk add --no-cache zip && cd /src/eb && zip -r /src/eb.zip .'

clean:
	rm -rf polyglot-co-demo-backend eb.zip
