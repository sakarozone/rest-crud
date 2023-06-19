GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
BINARY_NAME = restserver
DOCKER_IMAGE = mysql

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

docker-build:
	docker build -t $(DOCKER_IMAGE) .

installmockery:
	go install github.com/vektra/mockery/v2@v2.15.0

generatemocks:
    rm -rf mocks/* && \
    mockery --all --keeptree --inpackage

swagger: 
	swag init --dir ./cmd/movieserver --pd --parseDepth 10