deps: 
	go mod tidy && go install github.com/vektra/mockery/v2@v2.42.1

tests: 
	go test ./...  -cover -coverprofile=coverage.out

mocks: 
	mockery --dir=./src/domain --all --output=src/mocks

coverage: 
	go tool cover -html=coverage.out