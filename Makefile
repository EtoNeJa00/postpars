.PHONY: build
build: 
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Rest_Api_TV  

.DEFAULT_GOAL := build