ENV_VARS = CGO_ENABLED=0 GOOS=linux
FLAGS = -race
PACKAGE_NAME = api
INPUT_FILE = main.go
OUTPUT_FILE = ./tmp/main

.PHONY: build
build: 
	$(ENV_VARS) go build $(FLAGS) -o $(OUTPUT_FILE) $(INPUT_FILE)

.PHONY: run
run: build
	$(OUTPUT_FILE)

