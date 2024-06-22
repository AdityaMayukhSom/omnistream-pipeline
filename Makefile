ENV_VARS = CGO_ENABLED=0 GOOS=windows
FLAGS = -ldflags "-w -s"
PACKAGE_NAME = api
INPUT_FILE = main.go
OUTPUT_FILE = ./tmp/main.exe

.PHONY: build
build: build.docs
	$(ENV_VARS) go build $(FLAGS) -o $(OUTPUT_FILE) $(INPUT_FILE)

# we implicitely assume that `go install github.com/swaggo/swag/cmd/swag@latest` is installed
.PHONT: build.docs
build.docs:
	swag init

.PHONY: run
run: build
	$(OUTPUT_FILE)

