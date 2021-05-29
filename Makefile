APP_NAME=$(shell basename $(shell pwd))
APP_LOG=/var/log/${APP_NAME}.log
IMAGE_NAME=event-data-loader


.PHONY: \
		all \
		initgomod \
		protoc \
		unitest \
		test \
		clean \
		install \
		uninstall \
		start \
		stop \
		docker \
		push \
		reload

all: ./protoc ./${APP_NAME} unitest

initgomod:
	cd gosrc && go mod init ${APP_NAME} && go mod tidy

./${APP_NAME}: \
		./gosrc/cmd/main.go \
		./Makefile
	cd gosrc && go build -o ../${APP_NAME} ${APP_NAME}/cmd

unitest:
	# no tests yet
	#cd gosrc && go test ${APP_NAME}
	

test: ./protoc ./${APP_NAME} ./unitest

clean:
	rm -f ./${APP_NAME}