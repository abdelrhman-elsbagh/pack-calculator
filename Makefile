APP_NAME=pack-calculator-api
PORT?=8080

build:
	go build -o main ./cmd

run:
	go run ./cmd

test:
	go test ./...

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker-compose up --build

clean:
	rm -f main