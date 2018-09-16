GOFILES= $$(go list -f '{{join .GoFiles " "}}')

clean:
	rm -rf vendor/

deps: clean
	glide install

test:
	go test -timeout=5s -cover -race $$(glide novendor)

run:
	go run $(GOFILES) server -config="config.yaml"

build:
	go build -o $(GOPATH)/bin/prog-image $(GOFILES)

migrate:
	go run main.go -migrate=up

migrate_down:
	go run main.go -migrate=down

docker:
	docker build -t prog-image .
	docker run -d --name prog1 -p 8080:8080 prog-image
