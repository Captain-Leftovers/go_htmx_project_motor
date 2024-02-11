# Variables
BINARY_NAME := out
BUILD_DIR := tmp/

run: 
	@templ generate
	@go build -o ${BUILD_DIR}${BINARY_NAME} && ${BUILD_DIR}${BINARY_NAME}