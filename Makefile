mocks:
	cd ./store/mocks/; go generate;

build:
	go build -o cmd/api/simple-webserver cmd/api/main.go

run:
	cd cmd/api; ./rundev.sh

#	for create migration
#	migrate create -ext sql -dir store/mysql/migrations -seq create_users_table
