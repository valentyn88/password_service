# Parameters
MAIN_FILE=main.go
MAIN_PATH=cmd/password_server
SERVER=password_server

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o $(MAIN_PATH)/$(SERVER) $(MAIN_PATH)/$(MAIN_FILE)

build-linux:
	GOARCH=amd64 GOOS=linux go build -o $(MAIN_PATH)/$(SERVER) $(MAIN_PATH)/$(MAIN_FILE)

run:
	$(MAIN_PATH)/$(SERVER)

test:
	go test -race -cover -v ./...