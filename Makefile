mocks:
	cd ./store/mocks/; go generate;

build:
	go build -o cmd/api/simple-webserver cmd/api/main.go

run:
	cd cmd/api; ./rundev.sh