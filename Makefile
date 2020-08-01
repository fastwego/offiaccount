TIME := $(shell date '+%Y-%m-%d %H:%M:%S')
HASH := $(shell git describe --always)

build:
	go fmt ./...
	goimports -w ./
	addlicense -c FastWeGo -v -l apache ./*.go
	addlicense -c FastWeGo -v -l apache ./*/*.go
	addlicense -c FastWeGo -v -l apache ./*/*/*.go
	addlicense -c FastWeGo -v -l apache ./*/*/*/*.go

