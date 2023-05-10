build: 
	@go build -o bin/goBookings

run: build
	@./bin/goBookings

test:
	@go test -v ./...



