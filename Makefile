build:
	docker build --rm -t latex .
run:
	go run main.go | docker run -i latex > ot.pdf
example:
	curl -X POST -H "Content-Type: text/plain" --data "$(cat /Users/murr/go/src/resumebuilder/test/config.toml)" http://localhost:9001/api/toml 