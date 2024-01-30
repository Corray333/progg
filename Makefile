.SILENT:
build:
	cd cmd && go build main.go
front:
	cd frontend && npm run serve
run: build
	cd cmd && ./main