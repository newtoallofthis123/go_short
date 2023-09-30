build:
	@go build -o "bin/short.exe"

run: build
	@bin/short.exe