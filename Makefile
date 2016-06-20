.PHONY=clean

default: polyglot-co-demo-backend

polyglot-co-demo-backend:
	docker run --rm \
	  -v "$$(pwd):/src" \
	  -v /var/run/docker.sock:/var/run/docker.sock \
	  -e COMPRESS_BINARY=true \
	  centurylink/golang-builder

clean:
	rm -rf polyglot-co-demo-backend
