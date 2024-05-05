test:
	go test ./...

build:
	go build -o raytrace -ldflags="-s -w"