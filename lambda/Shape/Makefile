all: cover

LAMBDA_NAME := $(shell grep -A3 'artifactId:' lambda.yaml | tail -n1 | awk '{ print $$2}')
BINARY_NAME=bootstrap
ARCH=arm64
BUILD_FOLDER=target

hash:
	cd ${BUILD_FOLDER}; \
	openssl dgst -sha256 -binary ./${LAMBDA_NAME}.zip | openssl enc -base64 > ./${LAMBDA_NAME}.hash

zip:
	cd ${BUILD_FOLDER}; \
	zip ${LAMBDA_NAME}.zip ${BINARY_NAME}

build:
	mkdir -p ${BUILD_FOLDER}; \
	cd ${BUILD_FOLDER}; \
	GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build -o ${BINARY_NAME} ../cmd/main.go

update: build zip
	cd ${BUILD_FOLDER}; \
	aws lambda update-function-code --function-name ${LAMBDA_NAME} --zip-file fileb://${LAMBDA_NAME}.zip --region us-east-1 --profile $(profile)
