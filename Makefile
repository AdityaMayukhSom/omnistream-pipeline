FLAGS = 
PACKAGE_NAME = api
INPUT_DIR = main.go
OUTPUT_DIR = ./build/



.PHONY: clean
clean:
	rm -r ./build/*

.PHONY: build
build:
	go build $(FLAGS) -o $(OUTPUT_DIR) $(INPUT_DIR)

.PHONY: run
run: clean build
	./build/main.exe

