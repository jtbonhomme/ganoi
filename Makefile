.PHONY: help clean test run debug lint pprof
IMAGES_TAG = ${shell git describe --tags --match '[0-9]*\.[0-9]*\.[0-9]*' 2> /dev/null || echo 'latest'}
GIT_SHA1:=$(shell git rev-parse --short HEAD)
REPO=jtbonhomme/ganoi

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

lint: ## Execute Golangci-lint on the repo.
	golangci-lint -v --deadline 100s --skip-dirs docs run ./...

test: lint ## Go test the repo.
	go test ./... -cover -coverprofile coverage.out

run: ## Run the main program.
	go run cmd/ganoi/main.go -config-file ./config.yml

debug: ## Run the main program.
	go run cmd/ganoi/main.go -config-file ./config.yml -debug -optim

pprof: ## Run the main program with profiling.
	go run cmd/ganoi/main.go -debug -cpuprofile profile.prof

clean: ## Build the main program.
	rm -f ganoi_*.dump
	rm -f screenshot_*.png
	rm -f profile*.png
	rm -f profile.prof
	rm -f ganoi ganoi.wasm
	rm -f coverage.out cover.xml cover.html

build: ## Build the main program.
	go build -o ganoi cmd/ganoi/main.go

wasm: ## Build for Web Assembly distribution.
	GOOS=js GOARCH=wasm go build -o build/ganoi.wasm cmd/ganoi/main.go

serve: ## Serve Web Assembly build on localhost:8080.
	which wasmserve || (go install github.com/hajimehoshi/wasmserve@latest)
	@echo "Open http://localhost:8080"
	wasmserve cmd/ganoi/main.go

badge: lint ## Generate a coverage badge.
	which gopherbadger || (go get github.com/jpoles1/gopherbadger)
	gopherbadger

cover: test ## Measure the test coverage.
	which gocov || (go get -u github.com/axw/gocov/gocov)
	which gocov-xml || (go get -u github.com/AlekSi/gocov-xml)
	which gocov-html || (go get -u github.com/matm/gocov-html)
	gocov convert coverage.out | gocov-xml > cover.xml
	gocov convert coverage.out | gocov-html > cover.html
	open cover.html

login: ## Log in to docker hub registry.
	@docker login

build-docker: login ## Build docker images.
	docker build \
		-t ${REPO}:latest \
		-t ${REPO}:${IMAGES_TAG} \
		-t ${REPO}:${GIT_SHA1} \
		-f pkg/${SRV_NAME}/Dockerfile .
	docker push ${REPO}:latest
	docker push ${REPO}:${IMAGES_TAG}
	docker push ${REPO}:${GIT_SHA1}
