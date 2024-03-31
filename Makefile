FLAGS = -race
PACKAGE_NAME = api
INPUT_DIR = ./$(PACKAGE_NAME)/main.go
OUTPUT_DIR = ./build/

.PHONY: build
build:
	go build $(FLAGS) -o $(OUTPUT_DIR) $(INPUT_DIR)

.PHONY: clean
clean:
	rm -r ./build/*

