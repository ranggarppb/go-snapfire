EXEC_PATH=bin/go-snapfire

run: build
	@./${EXEC_PATH}

build:
	go build -o ${EXEC_PATH}