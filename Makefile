
.PHONY :build
build:
	go build -o bin/damon ./cmd/damon

.PHONY :run
run: build
	./bin/damon

.PHONY :install-osx
install-osx:
	cp ./bin/damon /usr/local/bin/damon

.PHONY :test
test:
	go test ./...

.PHONY :get
gen:
	go generate ./...
