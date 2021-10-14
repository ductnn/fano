BINARY_NAME=fano

build: GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go

run: ./${BINARY_NAME}

build_and_run: build run

clean: rm ${BINARY_NAME}-darwin

