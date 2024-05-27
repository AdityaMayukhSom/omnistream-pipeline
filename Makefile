ENV_VARS = CGO_ENABLED=1 GOOS=windows
FLAGS = -race -ldflags "-w -s"
PACKAGE_NAME = api
INPUT_FILE = main.go
OUTPUT_FILE = ./tmp/main.exe

.PHONY: build
build: 
	$(ENV_VARS) go build $(FLAGS) -o $(OUTPUT_FILE) $(INPUT_FILE)

.PHONY: run
run: build
	$(OUTPUT_FILE)

