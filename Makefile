validate:
	swagger validate ./swagger/swagger.yaml


serve:
	swagger serve -F=swagger ./swagger/swagger.yaml


all:
	swagger validate ./swagger/swagger.yaml
	swagger serve -F=swagger ./swagger/swagger.yaml

tests:
	go test -covermode=count ./...

mockgen: ## Run mockgen cli fro generate mocks
	mockgen \
		-destination=database/mock.go \
		-package database \
		team-project/database TicketRepository, UserCRUD, Model, TripRepository

perf-test:
	vegeta attack -targets=src/performance/test-login.txt -duration=54s -rate=200 | tee src/performance/results.bin | vegeta report
	cat src/performance/results.bin | vegeta report -reporter=plot > src/performance/plot.html


go-build:
	GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' o team-project


dc-build:
	docker-compose build
dc-up:
	docker-compose up &

migrate-up:
	src/bin/goose -dir ./migrations postgres "user=mccuwyhjuexqwp password=cd534c973a2f026e46ff57ea94cf9fc6fa29951d7cdcaa18b96415d06e6264dd host=ec2-54-247-70-127.eu-west-1.compute.amazonaws.com dbname=da2utgpo2vg4ca sslmode=disable" up-to 20190426030708

migrate-down:
	src/bin/goose -dir ./migrations postgres "user=mccuwyhjuexqwp password=cd534c973a2f026e46ff57ea94cf9fc6fa29951d7cdcaa18b96415d06e6264dd host=ec2-54-247-70-127.eu-west-1.compute.amazonaws.com dbname=da2utgpo2vg4ca sslmode=disable" down-to 20190426030708
