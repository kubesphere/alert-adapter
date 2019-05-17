generate: Makefile
	docker build -t adapter -f ./Dockerfile .

dev:
	rm -f adapter
	echo "Building binary..."
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -ldflags '-w' -o ./adapter cmd/adapter/main.go
	echo "Building images..."
	docker build -t adapter -f ./Dockerfile.dev .
	echo "Built successfully"

clean:
	rm -f adapter