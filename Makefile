PROJECTNAME = $(shell basename "$(PWD)")

BINDIRECTORY = bin

REVISIONBRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISIONCOMMIT = $(shell git rev-parse HEAD | head -c 8)
REVISIONDATE = $(shell date +%Y.%m.%d-%H:%M:%S)
REVISIONVERSION = $(shell git describe)
REVISION = $(REVISIONBRANCH)-$(REVISIONCOMMIT)-$(REVISIONDATE)

LDFLAGS = -ldflags="-X main.revision=$(REVISION) -X main.version=$(REVISIONVERSION)"

BUILDFLAGS = -v

build:
	go mod download && CGO_ENABLED=0 go build $(LDFLAGS) $(BUILDFLAGS) -o $(BINDIRECTORY)/anon-mail ./cmd/anon-mail/*.go

test:
	go test ./...

test.v:
	go test -v ./...

run: build
	$(BINDIRECTORY)/anon-mail

docker.run:
	docker-compose up -d anon-mail

docker.log:
	docker-compose logs -f

# docker.build:

# docker.deploy: docker.build

mock.generate:
	