BINARY_NAME=fano

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

clean:
	rm ${BINARY_NAME}

install:
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	sudo ln -s $(PWD)/$(BINARY_NAME) /usr/local/bin/

uninstall:
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
