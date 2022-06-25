gen-mock:
	mockgen -source=./internals/domains/repository/repository.go -destination=./internals/domains/repository/mocks/repository.go -package=mocks

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

run-http:
	go run main.go serveHttp
