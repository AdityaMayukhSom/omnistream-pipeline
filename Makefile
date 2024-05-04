FLAGS =
PACKAGE_NAME = api
INPUT_DIR = main.go
OUTPUT_DIR = ./build/

.PHONY: clean
clean:
	rm -r $(OUTPUT_DIR)*

.PHONY: build
build:
	go build $(FLAGS) -o $(OUTPUT_DIR) $(INPUT_DIR)

.PHONY: run
run: build
	./build/main.exe

