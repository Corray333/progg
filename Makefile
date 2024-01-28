.SILENT:
build:
	cd cmd && go build main.go
run: build test
	cd cmd && ./main
test:
	cd frontend && npm run serve
